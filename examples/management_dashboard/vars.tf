// These variables would commonly be defined as environment variables or sourced in a .env file
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

// value used below is JSON-escaped string of JSON present in sample.json file
variable "test_import_details" {
  default = "{\n  \"dashboards\": [\n    {\n      \"dashboardId\": \"ocid1.managementdashboard.dev..aaaaaaaazrrxainoaive7adj77uqejld45vch7zkoqrlh5fwv2_dummy_ocids\",\n      \"providerId\": \"log-analytics\",\n      \"providerName\": \"Logging Analytics\",\n      \"providerVersion\": \"2.0\",\n      \"tiles\": [\n        {\n          \"displayName\": \"Total Log Trend ($(nls.dashboard.TARGETS_WITH_STATUS))\",\n          \"savedSearchId\": \"ocid1.managementsavedsearch.dev..aaaaaaaakxkuh457mjwoenvzxfmuevgrt5levwhovolkp6fieq_dummy_ocids\",\n          \"row\": 0,\n          \"column\": 0,\n          \"height\": 3,\n          \"width\": 6,\n          \"nls\": {},\n          \"uiConfig\": {},\n          \"dataConfig\": [],\n          \"state\": \"DEFAULT\",\n          \"drilldownConfig\": []\n        },\n        {\n          \"displayName\": \"Simple table\",\n          \"savedSearchId\": \"ocid1.managementsavedsearch.dev..aaaaaaaaycxfhipu5wjd56jrzbnffz633nwzgtt4cvuejhk5vr_dummy_ocids\",\n          \"row\": 0,\n          \"column\": 6,\n          \"height\": 3,\n          \"width\": 6,\n          \"nls\": {},\n          \"uiConfig\": {\n            \"vizType\": \"table\",\n            \"defaultDataSource\": \"severityValues\",\n            \"tableInfo\": {\n              \"cols\": [\n                {\n                  \"val\": \"severity\"\n                },\n                {\n                  \"val\": \"value\"\n                }\n              ]\n            }\n          },\n          \"dataConfig\": [\n            {\n              \"name\": \"severityValues\",\n              \"type\": \"staticJsonDataSource\",\n              \"parameters\": {\n                \"jsonData\": [\n                  {\n                    \"severity\": \"CRITICAL\",\n                    \"value\": 1\n                  },\n                  {\n                    \"severity\": \"ERROR\",\n                    \"value\": 2\n                  },\n                  {\n                    \"severity\": \"WARNING\",\n                    \"value\": 5\n                  }\n                ]\n              }\n            }\n          ],\n          \"state\": \"DEFAULT\",\n          \"drilldownConfig\": []\n        },\n        {\n          \"displayName\": \"Simple Pie Chart\",\n          \"savedSearchId\": \"ocid1.managementsavedsearch.dev..aaaaaaaa2ir5taiwkghohoe2ajr4pjkn6vt56mh2remeeljjnp_dummy_ocids\",\n          \"row\": 3,\n          \"column\": 0,\n          \"height\": 3,\n          \"width\": 6,\n          \"nls\": {},\n          \"uiConfig\": {},\n          \"dataConfig\": [],\n          \"state\": \"DEFAULT\",\n          \"drilldownConfig\": []\n        }\n      ],\n      \"displayName\": \"Sample dashboard 1 - Dummy\",\n      \"description\": \"Dashboard that shows some sample widgets.\",\n      \"compartmentId\": \"ocid1.compartment.oc1..aaaaaaaa2spgr6ewrwzxqddbrv4y7ty6xiqzbh7n7bv3n6pd_dummy_ocids\",\n      \"isOobDashboard\": false,\n      \"isShowInHome\": false,\n      \"createdBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n      \"timeCreated\": \"2020-12-15T17:21:39.605Z\",\n      \"updatedBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n      \"timeUpdated\": \"2020-12-15T17:21:39.605Z\",\n      \"metadataVersion\": \"2.0\",\n      \"isShowDescription\": true,\n      \"screenImage\": \"....\",\n      \"nls\": {\n        \"TARGETS_WITH_STATUS\": {\n          \"key\": \"dashboardMetadata.agent_dashboard_100.DATA_SENT_TITLE\",\n          \"bundle\": \"resources/nls/DashboardMsg\"\n        }\n      },\n      \"uiConfig\": {\n        \"isFilteringEnabled\": true,\n        \"isTimeRangeEnabled\": true,\n        \"isRefreshEnabled\": true\n      },\n      \"dataConfig\": [],\n      \"type\": \"normal\",\n      \"isFavorite\": false,\n      \"savedSearches\": [\n        {\n          \"id\": \"ocid1.managementsavedsearch.dev..aaaaaaaaycxfhipu5wjd56jrzbnffz633nwzgtt4cvuejhk5vr_dummy_ocids\",\n          \"displayName\": \"Generic Chart\",\n          \"providerId\": \"log-analytics\",\n          \"providerVersion\": \"2.0\",\n          \"providerName\": \"Logging Analytics\",\n          \"compartmentId\": \"ocid1.compartment.oc1..aaaaaaaa2spgr6ewrwzxqddbrv4y7ty6xiqzbh7n7bv3n6pd_dummy_ocids\",\n          \"isOobSavedSearch\": false,\n          \"description\": \"This is a sample Widget to demonstrate saved widget metadata\",\n          \"nls\": {},\n          \"type\": \"WIDGET_SHOW_IN_DASHBOARD\",\n          \"uiConfig\": {},\n          \"dataConfig\": [],\n          \"createdBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"updatedBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"timeCreated\": \"2020-12-15T17:21:39.605Z\",\n          \"timeUpdated\": \"2020-12-15T17:21:39.605Z\",\n          \"screenImage\": \"...\",\n          \"metadataVersion\": \"2.0\",\n          \"widgetTemplate\": \"visualizations/chartWidgetTemplate.html\",\n          \"widgetVM\": \"visualizations/chartWidget\",\n          \"freeformTags\": {},\n          \"definedTags\": {}\n        },\n        {\n          \"id\": \"ocid1.managementsavedsearch.dev..aaaaaaaakxkuh457mjwoenvzxfmuevgrt5levwhovolkp6fieq_dummy_ocids\",\n          \"displayName\": \"Widget1\",\n          \"providerId\": \"log-analytics\",\n          \"providerVersion\": \"2.0\",\n          \"providerName\": \"Logging Analytics\",\n          \"compartmentId\": \"ocid1.compartment.oc1..aaaaaaaa2spgr6ewrwzxqddbrv4y7ty6xiqzbh7n7bv3n6pd_dummy_ocids\",\n          \"isOobSavedSearch\": false,\n          \"description\": \"This is a sample Widget to demonstrate saved widget metadata\",\n          \"nls\": {},\n          \"type\": \"WIDGET_SHOW_IN_DASHBOARD\",\n          \"uiConfig\": {\n            \"timeSelection\": \"Last 60 mins\",\n            \"entity\": \"qosd1b4joxnanbzk37w4\",\n            \"compartment\": \"OMC\",\n            \"option4\": \"others\"\n          },\n          \"dataConfig\": [],\n          \"createdBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"updatedBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"timeCreated\": \"2020-12-15T17:21:39.605Z\",\n          \"timeUpdated\": \"2020-12-15T17:21:39.605Z\",\n          \"screenImage\": \"...\",\n          \"metadataVersion\": \"2.0\",\n          \"widgetTemplate\": \"visualizations/widgetTemplate1.html\",\n          \"widgetVM\": \"visualizations/Widget1\",\n          \"freeformTags\": {},\n          \"definedTags\": {}\n        },\n        {\n          \"id\": \"ocid1.managementsavedsearch.dev..aaaaaaaa2ir5taiwkghohoe2ajr4pjkn6vt56mh2remeeljjnp_dummy_ocids\",\n          \"displayName\": \"Pie chart\",\n          \"providerId\": \"log-analytics\",\n          \"providerVersion\": \"2.0\",\n          \"providerName\": \"Logging Analytics\",\n          \"compartmentId\": \"ocid1.compartment.oc1..aaaaaaaa2spgr6ewrwzxqddbrv4y7ty6xiqzbh7n7bv3n6pd_dummy_ocids\",\n          \"isOobSavedSearch\": false,\n          \"description\": \"This is a sample Widget to demonstrate saved widget metadata\",\n          \"nls\": {},\n          \"type\": \"WIDGET_SHOW_IN_DASHBOARD\",\n          \"uiConfig\": {\n            \"defaultDataSource\": \"severityValues\",\n            \"chartInfo\": {\n              \"colorBy\": \"severity\",\n              \"value\": \"value\",\n              \"group\": \"\",\n              \"series\": \"severity\"\n            },\n            \"advancedChartSettings\": {\n              \"chartType\": \"pie\",\n              \"selectionMode\": \"none\",\n              \"orientation\": \"vertical\",\n              \"coordinateSystem\": \"cartesian\",\n              \"sorting\": \"off\",\n              \"stack\": \"off\",\n              \"stackLabel\": \"on\",\n              \"dataCursor\": \"on\",\n              \"legendRendered\": true,\n              \"legendPosition\": \"top\",\n              \"timeAxisType\": \"enabled\",\n              \"styleDefaults\": {\n                \"lineWidth\": 2\n              }\n            }\n          },\n          \"dataConfig\": [\n            {\n              \"name\": \"severityValues\",\n              \"type\": \"staticJsonDataSource\",\n              \"parameters\": {\n                \"jsonData\": [\n                  {\n                    \"severity\": \"CRITICAL\",\n                    \"value\": 1\n                  },\n                  {\n                    \"severity\": \"ERROR\",\n                    \"value\": 2\n                  },\n                  {\n                    \"severity\": \"WARNING\",\n                    \"value\": 5\n                  }\n                ]\n              }\n            }\n          ],\n          \"createdBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"updatedBy\": \"ocid1.compartment.oc1..aaaaaaaagbgan4ik55kburvwtgtbzaghmmz6jvf2qzsncgaws7_dummy_ocids\",\n          \"timeCreated\": \"2020-12-15T17:21:39.605Z\",\n          \"timeUpdated\": \"2020-12-15T17:21:39.605Z\",\n          \"screenImage\": \"...\",\n          \"metadataVersion\": \"2.0\",\n          \"widgetTemplate\": \"visualizations/chartWidgetTemplate.html\",\n          \"widgetVM\": \"visualizations/chartWidget\",\n          \"freeformTags\": {},\n          \"definedTags\": {}\n        }\n      ],\n      \"freeformTags\": {},\n      \"definedTags\": {}\n    }\n  ]\n}"
}