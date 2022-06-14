package testhelper

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"log"
	"strings"
)

const (
	IdentifyPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "identify",
	  "eventOrderNo": "1",
	  "context": {
		"traits": {
		  "trait1": "new-val"
		}
	  },
	  "timestamp": "2020-02-02T00:23:09.544Z"
	}`
	TrackPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "track",
	  "event": "%s",
	  "properties": {
		"review_id": "12345",
		"product_id": "123",
		"rating": 3,
		"review_body": "Average product, expected much more."
	  }
	}`
	PagePayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "page",
	  "name": "Home",
	  "properties": {
		"title": "Home | RudderStack",
		"url": "http://www.rudderstack.com"
	  }
	}`
	ScreenPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "screen",
	  "name": "Main",
	  "properties": {
		"prop_key": "prop_value"
	  }
	}`
	AliasPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "alias",
	  "previousId": "name@surname.com"
	}`
	GroupPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "group",
	  "groupId": "groupId",
	  "traits": {
		"name": "MyGroup",
		"industry": "IT",
		"employees": 450,
		"plan": "basic"
	  }
	}`
	ModifiedIdentifyPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "identify",
	  "context": {
		"traits": {
		  "trait1": "new-val"
		},
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  },
	  "timestamp": "2020-02-02T00:23:09.544Z"
	}`
	ModifiedTrackPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "track",
	  "event": "%s",
	  "properties": {
		"review_id": "12345",
		"product_id": "123",
		"rating": 3,
		"revenue": 4.99,
		"review_body": "Average product, expected much more."
	  },
	  "context": {
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  }
	}`
	ModifiedPagePayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "page",
	  "name": "Home",
	  "properties": {
		"title": "Home | RudderStack",
		"url": "http://www.rudderstack.com"
	  },
	  "context": {
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  }
	}`
	ModifiedScreenPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "screen",
	  "name": "Main",
	  "properties": {
		"prop_key": "prop_value"
	  },
	  "context": {
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  }
	}`
	ModifiedAliasPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "alias",
	  "previousId": "name@surname.com",
	  "context": {
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  }
	}`
	ModifiedGroupPayload = `{
	  "userId": "%s",
	  "messageId": "%s",
	  "type": "group",
	  "groupId": "groupId",
	  "traits": {
		"name": "MyGroup",
		"industry": "IT",
		"employees": 450,
		"plan": "basic"
	  },
	  "context": {
		"ip": "14.5.67.21",
		"library": {
		  "name": "http"
		}
	  }
	}`
)

func JsonEscape(i string) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", fmt.Errorf("could not escape big query JSON credentials for workspace config with error: %s", err.Error())
	}
	return strings.Trim(string(b), `"`), nil
}

func ConnectWithBackoff(operation func() error) {
	var err error

	backoffWithMaxRetry := backoff.WithMaxRetries(backoff.NewConstantBackOff(ConnectBackoffDuration), uint64(ConnectBackoffRetryMax))
	if err = backoff.Retry(operation, backoffWithMaxRetry); err != nil {
		log.Panicf("could not connect to warehouse with error: %s", err.Error())
	}
}

func GWJobsForUserIdWriteKey() string {
	return `CREATE OR REPLACE FUNCTION gw_jobs_for_user_id_and_write_key(user_id varchar, write_key varchar)
								RETURNS TABLE
										(
											job_id varchar
										)
							AS
							$$
							DECLARE
								table_record RECORD;
								batch_record jsonb;
							BEGIN
								FOR table_record IN SELECT * FROM gw_jobs_1 where (event_payload ->> 'writeKey') = write_key
									LOOP
										FOR batch_record IN SELECT * FROM jsonb_array_elements((table_record.event_payload ->> 'batch')::jsonb)
											LOOP
												if batch_record ->> 'userId' != user_id THEN
													CONTINUE;
												END IF;
												job_id := table_record.job_id;
												RETURN NEXT;
												EXIT;
											END LOOP;
									END LOOP;
							END;
							$$ LANGUAGE plpgsql`
}

func BRTJobsForUserId() string {
	return `CREATE OR REPLACE FUNCTION brt_jobs_for_user_id(user_id varchar)
									RETURNS TABLE
											(
												job_id varchar
											)
								AS
								$$
								DECLARE
									table_record  RECORD;
									event_payload jsonb;
								BEGIN
									FOR table_record IN SELECT * FROM batch_rt_jobs_1
										LOOP
											event_payload = (table_record.event_payload ->> 'data')::jsonb;
											if event_payload ->> 'user_id' = user_id Or event_payload ->> 'id' = user_id Or event_payload ->> 'USER_ID' = user_id Or event_payload ->> 'ID' = user_id THEN
												job_id := table_record.job_id;
												RETURN NEXT;
											END IF;
										END LOOP;
								END ;
								$$ LANGUAGE plpgsql`
}

func DefaultEventMap() EventsCountMap {
	return EventsCountMap{
		"identifies": 1,
		"users":      1,
		"tracks":     1,
		"pages":      1,
		"screens":    1,
		"aliases":    1,
		"groups":     1,
		"gateway":    6,
		"batchRT":    8,
	}
}