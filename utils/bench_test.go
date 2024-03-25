package utils_test

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"goLibrary/config"
	"testing"
)

type RelationSuite struct {
	suite.Suite
}

func (s *RelationSuite) SetupTest() {
}

func BenchmarkStr(b *testing.B) {
	var confi *config.Config
	for i := 0; i < b.N; i++ {
		confi = &config.Config{}
	}
	_ = confi
}

func BenchmarkRefectNew(b *testing.B) {
	var confi interface{}
	for i := 0; i < b.N; i++ {
		confi = &config.Config{}
		confi = confi.(*config.Config)
	}
	_ = confi
}

type T1 struct {
	A      int
	B      string
	Output T
}

type T struct {
	PersonId   string `json:"person_id"`
	Date       string `json:"date"`
	BioMetrics struct {
		SleepQuality   int `json:"sleep_quality"`
		StepsCount     int `json:"steps_count"`
		CaloriesBurned int `json:"calories_burned"`
		HeartRate      struct {
			Average int `json:"average"`
			Min     int `json:"min"`
			Max     int `json:"max"`
		} `json:"heart_rate"`
		BloodPressure struct {
			Systolic  int `json:"systolic"`
			Diastolic int `json:"diastolic"`
		} `json:"blood_pressure"`
		HydrationLevel struct {
			TotalIntakeMilliliters int     `json:"total_intake_milliliters"`
			WaterBalance           float64 `json:"water_balance"`
		} `json:"hydration_level"`
		Nutrition struct {
			TotalCalories           int `json:"total_calories"`
			CarbohydratesPercentage int `json:"carbohydrates_percentage"`
			ProteinsPercentage      int `json:"proteins_percentage"`
			FatsPercentage          int `json:"fats_percentage"`
		} `json:"nutrition"`
	} `json:"bio_metrics"`
	MoodTracker struct {
		MorningMood   string `json:"morning_mood"`
		AfternoonMood string `json:"afternoon_mood"`
		EveningMood   string `json:"evening_mood"`
		StressLevel   string `json:"stress_level"`
	} `json:"mood_tracker"`
	ActivityLog []struct {
		Time            string `json:"time"`
		Activity        string `json:"activity"`
		Location        string `json:"location,omitempty"`
		DurationMinutes int    `json:"duration_minutes"`
		Food            string `json:"food,omitempty"`
		Description     string `json:"description,omitempty"`
		WorkoutType     string `json:"workout_type,omitempty"`
		Genre           string `json:"genre,omitempty"`
	} `json:"activity_log"`
	EnvironmentalFactors struct {
		Weather struct {
			Temperature      int    `json:"temperature"`
			WeatherCondition string `json:"weather_condition"`
		} `json:"weather"`
		AirQualityIndex  int      `json:"air_quality_index"`
		NoiseLevelDb     int      `json:"noise_level_db"`
		LocationsVisited []string `json:"locations_visited"`
	} `json:"environmental_factors"`
	DigitalUsage struct {
		ScreenTimeHours        float64 `json:"screen_time_hours"`
		SocialMediaTimeMinutes int     `json:"social_media_time_minutes"`
		EmailsSent             int     `json:"emails_sent"`
		CallsMade              int     `json:"calls_made"`
		MessagesSent           int     `json:"messages_sent"`
	} `json:"digital_usage"`
	PersonalNotes struct {
		Gratitude            string   `json:"gratitude"`
		Accomplishments      []string `json:"accomplishments"`
		SelfImprovementGoals struct {
			MeditationMinutes int `json:"meditation_minutes"`
			ReadingPages      int `json:"reading_pages"`
		} `json:"self_improvement_goals"`
	} `json:"personal_notes"`
}

// mysql 常用场合
// BenchmarkJson-24           78220             15295 ns/op
func BenchmarkJson(b *testing.B) {
	var data = []byte("{\"output\":{\n  \"person_id\": \"123456\",\n  \"date\": \"2024-03-25\",\n  \"bio_metrics\": {\n    \"sleep_quality\": 85,\n    \"steps_count\": 7500,\n    \"calories_burned\": 2200,\n    \"heart_rate\": {\n      \"average\": 70,\n      \"min\": 50,\n      \"max\": 120\n    },\n    \"blood_pressure\": {\n      \"systolic\": 120,\n      \"diastolic\": 80\n    },\n    \"hydration_level\": {\n      \"total_intake_milliliters\": 2500,\n      \"water_balance\": 0.8\n    },\n    \"nutrition\": {\n      \"total_calories\": 1800,\n      \"carbohydrates_percentage\": 50,\n      \"proteins_percentage\": 30,\n      \"fats_percentage\": 20\n    }\n  },\n  \"mood_tracker\": {\n    \"morning_mood\": \"energetic\",\n    \"afternoon_mood\": \"focused\",\n    \"evening_mood\": \"relaxed\",\n    \"stress_level\": \"low\"\n  },\n  \"activity_log\": [\n    {\n      \"time\": \"07:00\",\n      \"activity\": \"jogging\",\n      \"location\": \"park\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"09:00\",\n      \"activity\": \"breakfast\",\n      \"food\": \"oatmeal and fruits\",\n      \"duration_minutes\": 15\n    },\n    {\n      \"time\": \"11:00\",\n      \"activity\": \"meeting\",\n      \"description\": \"project planning\",\n      \"duration_minutes\": 60\n    },\n    {\n      \"time\": \"13:00\",\n      \"activity\": \"lunch\",\n      \"food\": \"chicken salad\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"15:00\",\n      \"activity\": \"gym\",\n      \"workout_type\": \"weight lifting\",\n      \"duration_minutes\": 45\n    },\n    {\n      \"time\": \"19:00\",\n      \"activity\": \"dinner\",\n      \"food\": \"grilled salmon with vegetables\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"21:00\",\n      \"activity\": \"reading\",\n      \"genre\": \"science fiction\",\n      \"duration_minutes\": 60\n    },\n    {\n      \"time\": \"23:00\",\n      \"activity\": \"sleep\",\n      \"duration_minutes\": 480\n    }\n  ],\n  \"environmental_factors\": {\n    \"weather\": {\n      \"temperature\": 22,\n      \"weather_condition\": \"sunny\"\n    },\n    \"air_quality_index\": 30,\n    \"noise_level_db\": 35,\n    \"locations_visited\": [\n      \"home\",\n      \"park\",\n      \"office\",\n      \"gym\",\n      \"restaurant\"\n    ]\n  },\n  \"digital_usage\": {\n    \"screen_time_hours\": 4.5,\n    \"social_media_time_minutes\": 90,\n    \"emails_sent\": 20,\n    \"calls_made\": 5,\n    \"messages_sent\": 40\n  },\n  \"personal_notes\": {\n    \"gratitude\": \"Had a great day at work and a nice evening run.\",\n    \"accomplishments\": [\n      \"Completed project presentation\",\n      \"Cooked a healthy dinner\"\n    ],\n    \"self_improvement_goals\": {\n      \"meditation_minutes\": 15,\n      \"reading_pages\": 50\n    }\n  }\n}}")
	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &T1{})
	}
}

// BenchmarkJson2-24          78206             15178 ns/op
func BenchmarkJson2(b *testing.B) {
	var data = []byte("{\n  \"person_id\": \"123456\",\n  \"date\": \"2024-03-25\",\n  \"bio_metrics\": {\n    \"sleep_quality\": 85,\n    \"steps_count\": 7500,\n    \"calories_burned\": 2200,\n    \"heart_rate\": {\n      \"average\": 70,\n      \"min\": 50,\n      \"max\": 120\n    },\n    \"blood_pressure\": {\n      \"systolic\": 120,\n      \"diastolic\": 80\n    },\n    \"hydration_level\": {\n      \"total_intake_milliliters\": 2500,\n      \"water_balance\": 0.8\n    },\n    \"nutrition\": {\n      \"total_calories\": 1800,\n      \"carbohydrates_percentage\": 50,\n      \"proteins_percentage\": 30,\n      \"fats_percentage\": 20\n    }\n  },\n  \"mood_tracker\": {\n    \"morning_mood\": \"energetic\",\n    \"afternoon_mood\": \"focused\",\n    \"evening_mood\": \"relaxed\",\n    \"stress_level\": \"low\"\n  },\n  \"activity_log\": [\n    {\n      \"time\": \"07:00\",\n      \"activity\": \"jogging\",\n      \"location\": \"park\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"09:00\",\n      \"activity\": \"breakfast\",\n      \"food\": \"oatmeal and fruits\",\n      \"duration_minutes\": 15\n    },\n    {\n      \"time\": \"11:00\",\n      \"activity\": \"meeting\",\n      \"description\": \"project planning\",\n      \"duration_minutes\": 60\n    },\n    {\n      \"time\": \"13:00\",\n      \"activity\": \"lunch\",\n      \"food\": \"chicken salad\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"15:00\",\n      \"activity\": \"gym\",\n      \"workout_type\": \"weight lifting\",\n      \"duration_minutes\": 45\n    },\n    {\n      \"time\": \"19:00\",\n      \"activity\": \"dinner\",\n      \"food\": \"grilled salmon with vegetables\",\n      \"duration_minutes\": 30\n    },\n    {\n      \"time\": \"21:00\",\n      \"activity\": \"reading\",\n      \"genre\": \"science fiction\",\n      \"duration_minutes\": 60\n    },\n    {\n      \"time\": \"23:00\",\n      \"activity\": \"sleep\",\n      \"duration_minutes\": 480\n    }\n  ],\n  \"environmental_factors\": {\n    \"weather\": {\n      \"temperature\": 22,\n      \"weather_condition\": \"sunny\"\n    },\n    \"air_quality_index\": 30,\n    \"noise_level_db\": 35,\n    \"locations_visited\": [\n      \"home\",\n      \"park\",\n      \"office\",\n      \"gym\",\n      \"restaurant\"\n    ]\n  },\n  \"digital_usage\": {\n    \"screen_time_hours\": 4.5,\n    \"social_media_time_minutes\": 90,\n    \"emails_sent\": 20,\n    \"calls_made\": 5,\n    \"messages_sent\": 40\n  },\n  \"personal_notes\": {\n    \"gratitude\": \"Had a great day at work and a nice evening run.\",\n    \"accomplishments\": [\n      \"Completed project presentation\",\n      \"Cooked a healthy dinner\"\n    ],\n    \"self_improvement_goals\": {\n      \"meditation_minutes\": 15,\n      \"reading_pages\": 50\n    }\n  }\n}")
	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &T{})
	}
}
func TestConvBench(t *testing.T) {
	suite.Run(t, new(RelationSuite))

}
