// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package jwt includes sample JWT Token used in e2e tests.
package jwt

const (
	// Payload {
	//  "exp": 4715782722,
	//  "groups": [
	//    "group-1"
	//  ],
	//  "iat": 1562182722,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --listclaim groups "group-1"
	// nolint: lll
	TokenIssuer1 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3MTU3ODI3MjIsImdyb3VwcyI6WyJncm91cC0xIl0sImlhdCI6MTU2MjE4MjcyMiwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.JUWnHmKP01ZJMxtE8ke7Lv7G_eXL9yEoa4xWi__Sk09IQsuBGnhK75nQfl4aJ1fLthu-AvIp-GZe-1hDnBHhO4MQ7tXYeCdgnFkVGvhvZwOWpOl79H2U-mRX1eyjsHG0ZzwHwWMMLpPW6HQYCbTCEkZKRtoFoSum0-ilGqLO7EGRMEFpKYKnZV0N9pZIgX8uPlHoOVCZjqvhDtFjP-PG89vvNozeV7u47LMMzQeQtzCDmbRrIw1KVHZ885N1qiLcf-RNR9x9UINK_d98QKqkYsJtI2QTdVzevzSF8LfvaGLh75tvAL8h7L221haAKEnICuRqGWNO8qIwsj0k89kdGA"

	// Payload {
	//  "aud": foo,
	//  "exp": 4732994801,
	//  "iat": 1579394801,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --aud=foo
	// nolint: lll
	TokenIssuer1WithAud = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJmb28iLCJleHAiOjQ3MzI5OTQ4MDEsImlhdCI6MTU3OTM5NDgwMSwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.OmCSy8PplBCwdLnn_uOy3ADq_9bjXznyNwhyuCI5uaaZzWusHNMg0KKEk-2vKo9adbZKHfqGk0ID3ONWOPIy2WYkpHTMU_4FY7dP7pdORjqIVuIQ6ZNLVLjiOg_LM5Se-Gvq53p61Zqb25Io_gyHqsXzyYDrWOzYDmkjOK6HNjPLhEgnZipwZ2eYiZnhYXwyYSnnuxV6FosbCvA-nFiJcbIsQObzBGC6SDZhtuAcZf1i54K16se1yiiI-LwUkyA3DZhf0P-pWEo-LfBCMVqHBuhBGeQ6twfbfU5O5KRxroxno8v7uJON0LFqP-4uv1CX2Fta4MT_hAizqNX1z4Bkxw"

	// Payload {
	//  "azp": "bar",
	//  "exp": 4734125453,
	//  "iat": 1580525453,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --claims "azp:bar"
	// nolint: lll
	TokenIssuer1WithAzp = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJhenAiOiJiYXIiLCJleHAiOjQ3MzQxMjU0NTMsImlhdCI6MTU4MDUyNTQ1MywiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.SO4qjRJPYItkpGGpCDfEhaUfdthO8L9b_aawao4dJKyqqXN0uYdsJau_JZzyPQ1emAmJP7VyjwELrlszA6xV65na_O-eny23iwhEoroChQMpcr9DWqSUBUfpbHSPFAjUv38SUbQfLgar0HrMxQlTAzB0vyzn2g6-cukP469ZlOUmzvi9b4UpolTLp_WPgEHKjZw8CL56CcuJqBIKgfn0M7ta2bY_qx-UrsEW0CqnXol7vhXuDAfMeWZYKuDP8qc2VH1T6wpO2JnZ0EaNDuZfQLOWFYKsFGlaYcus9j462AfJQBSFQTbkIjkvKMK6aI_rMEesAnJr2eei1UYi14JYiQ"

	// Payload {
	//  "exp": 4757607896,
	//  "iat": 1604007896,
	//  "iss": "test-issuer-1@istio.io",
	//  "nested": {
	//    "key1": [
	//      "valueA",
	//      "valueB"
	//    ],
	//    "nested-2": {
	//      "key1": [
	//        "valueA",
	//        "valueB"
	//      ]
	//    }
	//  },
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --nestedclaim key1 valueA valueB
	// nolint: lll
	TokenIssuer1WithNestedClaims1 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3NTc2MDc4OTYsImlhdCI6MTYwNDAwNzg5NiwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsIm5lc3RlZCI6eyJrZXkxIjpbInZhbHVlQSIsInZhbHVlQiJdLCJuZXN0ZWQtMiI6eyJrZXkxIjpbInZhbHVlQSIsInZhbHVlQiJdfX0sInN1YiI6InN1Yi0xIn0.l8p-9Yha6uAEC4iahFNHvMxeEvQmBYxegGCiBr2oJpsmKu4IxuOL9RnEQpG5C2xHgF9vBfjifo_Jl1qtMCiKMRsWoFusnNoPwbk3mUp1Y8t1mvOkVPp4roed6tM5NmhUQMyP106sru4DszS1QNfDyGtpZZny2ls5Mva-B3kxclLjswXuRb5v70YKfqpYVDtZbltZg5CD2-V6ZwJgx2ZlXCOob-7uCDyaLSRWLxS8tFMo_akdWFCxpnkGJUtS5R1nAcofX5vYSqxfcKMSJpKMluDcX859jQ9zu_eBfAjvniVP8iOuZU-JKI2GywL3x1hBfD4ebSvNiEs7jCszagIM8Q"

	// Payload {
	//  "exp": 4757608018,
	//  "iat": 1604008018,
	//  "iss": "test-issuer-1@istio.io",
	//  "nested": {
	//    "key2": "valueC",
	//    "nested-2": {
	//      "key2": "valueC"
	//    }
	//  },
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --nestedclaim key2 valueC
	// nolint: lll
	TokenIssuer1WithNestedClaims2 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3NTc2MDgwMTgsImlhdCI6MTYwNDAwODAxOCwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsIm5lc3RlZCI6eyJrZXkyIjoidmFsdWVDIiwibmVzdGVkLTIiOnsia2V5MiI6InZhbHVlQyJ9fSwic3ViIjoic3ViLTEifQ.dPBVaeOoMBFh9EFovHtmG2V-Bw106kUAV4GTPTVwoOaLxVUfAIpumrGX1fvEYLmtdwEkkL2CkkDdmKbsBgJBfnwg04oMeyHHrpSe5Y-8NJWGGxaBTzc7btWIbc1IPvj0vw_jIBeSqpXyW2Kd1Rc9kxK8rRuTdAeSK6oKcqg5d6kw-fC5PeiDlnuBH7F-xIaPYOHSxlKjC6TSP6-JYlQHw-uZ6rJmFwhUvc9B06dl3Xgx1W3PmYyt1e4kmHBPbwTgj1ckv8yev1JfIbciaNyQcgMf_eFwgQscHSeUY2e3v6DYAGHAQv0M6o_wweUlHfJP1juxUVA3UfpqC76NeiYipg"

	// TokenIssuer1WithCollisionResistantName
	// payload:
	//{
	//  "exp": 4837231657,
	//  "iat": 1683631657,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1",
	//  "test-issuer-1@istio.io/nested": {
	//    "key1": "valueC",
	//    "nested-2": {
	//      "key1": "valueC"
	//    }
	//  },
	//  "test-issuer-1@istio.io/simple": "valueC"
	//}
	// Generated by: python3 security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1
	// --claims=test-issuer-1@istio.io/simple:valueC
	// --nestedkey=test-issuer-1@istio.io/nested --nestedclaim key1 valueC
	// nolint: lll
	TokenIssuer1WithCollisionResistantName = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ4MzcyMzE2NTcsImlhdCI6MTY4MzYzMTY1NywiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIiwidGVzdC1pc3N1ZXItMUBpc3Rpby5pby9uZXN0ZWQiOnsia2V5MSI6InZhbHVlQyIsIm5lc3RlZC0yIjp7ImtleTEiOiJ2YWx1ZUMifX0sInRlc3QtaXNzdWVyLTFAaXN0aW8uaW8vc2ltcGxlIjoidmFsdWVDIn0.kiXEekepka92nbKNZoPrjx4oeg5OI_ILROfZ9RR02tPsaXmVAEXzfsd5foend_seEru85oLXwiBM8uFJgn8hjE9oD4wXs-xGJHpgTogBeiuzIkrWN0o6iD03tYHFKCTUY8O5B5-9NoyYiEL5pNUqPd16VwfMRg_8qaOZPBVqES-mQED8xVL_UXOlkVQEsb2QUZD_9WfkzO1JRk8m1b6gxxxN2Its1jUB-F667yUQ5tUcYbuaus5yG5sWwpAJTXWzakC2Qq6NfB5hV2fwOsVSLUvH6B7sBq6dz6W03-xFH58j_aAAF0ceJ_LsX63hqZfUOeIcbKCRFoe6Ou8cw5sRtw"

	// Payload {
	//  "exp": 4715782783,
	//  "groups": [
	//    "group-2"
	//  ],
	//  "iat": 1562182783,
	//  "iss": "test-issuer-2@istio.io",
	//  "sub": "sub-2"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-2@istio.io --sub=sub-2 --listclaim groups "group-2"
	// nolint: lll
	TokenIssuer2 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3MTU3ODI3ODMsImdyb3VwcyI6WyJncm91cC0yIl0sImlhdCI6MTU2MjE4Mjc4MywiaXNzIjoidGVzdC1pc3N1ZXItMkBpc3Rpby5pbyIsInN1YiI6InN1Yi0yIn0.KmOIdRoI-FfaWR6xdMWrguOypmqD5UC2VR6B1Cw8V0aBhaqpCBcn5fDvA13mSiU3sx3VFF7-CbnnZVz4j6k2RsYNnWDDJTzr4GnxIXJYJpIRBcE0Au9mcZM0cVhW9BBtKtc0twy0I5S7H7qIAEPBq6iYm8IED409K7YIuI92kLeSe--ehhs1yWmI-YbixYwwq7viD1ZIRdbP5b2ZxU6A0VHnWVXFBGSoYshSLdJS6fnlU480SXyz0oUakDpkaVkz6n9LXolggs6Cz6Gq9OIEs2zUz2DqAJ5h6HChjIQ8AydwQnRMx4qnEbdtqq2qy1h5HHfosq8pcGbiQKf7CSQKHw"

	// Payload {
	//  "exp": 4730732495,
	//  "iat": 1577132495,
	//  "iss": "test-issuer-2@istio.io",
	//  "permission": "read write",
	//  "sub": "sub-2"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-2@istio.io --sub=sub-2 --claims "permission:read write"
	// nolint: lll
	TokenIssuer2WithSpaceDelimitedScope = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3MzA3MzI0OTUsImlhdCI6MTU3NzEzMjQ5NSwiaXNzIjoidGVzdC1pc3N1ZXItMkBpc3Rpby5pbyIsInBlcm1pc3Npb24iOiJyZWFkIHdyaXRlIiwic3ViIjoic3ViLTIifQ.JvzbnDfcA2XD8oeX1kHacGKRCPoKH19Dv8LwJXfob4gbvzplBLk09tKDVOClJ0kFNC2S8mFJJVm-ASAqxJoOfrd74RhsKVtcKVYA7mk3YAnw8oXNrEDJ_sGLBvbVnIOOdylM1V6whoJjn_f6rSvSplzdxNXUQYbuzTn5iyGkA-YSD8zvvCaTkFQqSXBnhLJ-VWHvHMsFUrrcMHSwuGYPm5CblZrCK0ICDqZcKRZh1qw6Elw0ZE46qkZWryOP_e2ydenbgv8MxR7WXkN8zk0Q6XRVmKjQ83ndjULMjye2FFnklFdDbaPC3rR8x0vWgf92hB49IytpLGa95uacryNE1g"

	// Payload {
	//  "exp": 4847801434,
	//  "iat": 1694201434,
	//  "iss": "test-issuer-3istio.io",
	//  "permission": "read write",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-3@istio.io --sub=sub-1 --claims "permission:read write"
	// nolint: lll
	TokenIssuer3 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ4NDc4MDE0MzQsImdyb3VwcyI6WyJncm91cC0xIl0sImlhdCI6MTY5NDIwMTQzNCwiaXNzIjoidGVzdC1pc3N1ZXItM0Bpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.YaRHh-0yccSHinMxo5eseP_ySSLE3pnFt7oL4HHL7R3RuYfzQsoow1Pc96uShVfOifIFBiC2_sfGje6lw-cJDo-Lw1g3kQOdmAjoQjmuUJMXsRCuR4kxLtREPW-ifcMHrKGoErJrzU8FIiwXwzNBPvrnY4HDHlG40E5wHSyf19sg11NtudLQubjCVaFUwuPCgL6JUqA2y80KReI4lqEz_-G9YoMFCWPnlHK31GHSvuoO0PzJm-H9OF7PSLGZoff0Hs7m5IN2qTo1Mc_wK1QBEmJRHZv6DXNMUtPOAFYzayv3ND6KXzqN4KrAk3d7KW_pVAIBb0oiAiI3TZy2FjYO8g"

	// Payload {
	//  "exp": 1562182856,
	//  "groups": [
	//    "group-1"
	//  ],
	//  "iat": 1562182855,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=1 --iss=test-issuer-1@istio.io --sub=sub-1 --listclaim groups "group-1"
	// nolint: lll
	TokenExpired = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE1NjIxODI4NTYsImdyb3VwcyI6WyJncm91cC0xIl0sImlhdCI6MTU2MjE4Mjg1NSwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.o9hp0P5DS66Q7wP38wGo0AKS5HoWdHXrAUdNLnXzVC4MwU4N9o3U0PfDgWp8A7YIbDuuQAtKJyCHEJCV3JEh7Xb5Tz_12hYQcXcJ0FTA6pOXrbWRjkAVMhs3-OHiKt855s39l2OKrLwmJ3ph7LV4z8J8i-2LE2hQH18R00_dcx2agoY1VNYH2LSErBYx6e-HQlKFW8c9sQ1CHG7u4ns1I2e23A0nBlRWRUHUYIMo6sfAPWhyQWl1kpRzg6b3fyXGfUpgeEmdIPDK7MfRUZA-0epFGjvoqCwgMdEEQ-O_pH5y2qV1jPpu-9IO_FdpYhHofKMTn_ql05ys6zoIHj9Gng"

	TokenInvalid = "TokenInvalid"

	// Payload {
	//  "aud": [
	//	  "foo",
	// 	  "buzz"
	//  ],
	//  "exp": 4732994801,
	//  "iat": 1579394801,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --aud=foo,buzz
	// nolint: lll
	TokenIssuer1WithAudList = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJhdWQiOlsiZm9vIiwiYnV6eiJdLCJleHAiOjQ4Njg3MjU5OTMsImlhdCI6MTcxNTEyNTk5MywiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.A0rD0DeHInl2thoxfx69A3o8BQvFX_cHNS4-1KNkUsHCD9muyqj4eO0AdRY__wAPGbkzXbESbU5NpHwMXWa1zg6KOMHiLmtze4RmQ8XH3mFMi4eI-Qr2N11Ce40c9SvZurj5lWcRQCtOp4DBNtLvFhueqvpW6odRmuREaN11yxgpQeBM8EzhnF_njDCCqcai0FAms97Nx-fbXS2ASErPm73rmbuUpc5Usm3kzP22t3OP_Q8DlXO53xQPfaBCIo-FGZCWS1UzD37cplHGzKndK1V7PelaFnJE7U95yBdJaMufk1KSX_aPlF6Kjn4lABydlzg8y3IZh5ERZIgSX9cs4w"
)
