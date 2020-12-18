package main

type AdProperty struct {
	ExternalReference    string `json:"externalReference"`
	LastModificationDate string `json:"lastModificationDate"`
	ID                   int    `json:"id"`
	SEOURL               string `json:"SEOUrl"`
	Media                struct {
		Pictures struct {
			BaseURL string `json:"baseUrl"`
			Count   int    `json:"count"`
			Items   []struct {
				Orientation string `json:"orientation"`
				RelativeURL struct {
					Large  string `json:"large"`
					Medium string `json:"medium"`
				} `json:"relativeUrl"`
			} `json:"items"`
		} `json:"pictures"`
	} `json:"media"`
	Property struct {
		Basement struct {
			Surface int  `json:"surface"`
			Exists  bool `json:"exists"`
		} `json:"basement"`
		Parking struct {
			ParkingSpaceCount struct {
				Total  int `json:"total"`
				Indoor int `json:"indoor"`
			} `json:"parkingSpaceCount"`
		} `json:"parking"`
		Specificities struct {
			SME struct {
				Office struct {
					Surface int  `json:"surface"`
					Exists  bool `json:"exists"`
				} `json:"office"`
			} `json:"SME"`
		} `json:"specificities"`
		Type       string `json:"type"`
		DiningRoom struct {
			Exists bool `json:"exists"`
		} `json:"diningRoom"`
		Building struct {
			Condition         string  `json:"condition"`
			ConstructionYear  int     `json:"constructionYear"`
			FloorCount        int     `json:"floorCount"`
			FacadeCount       int     `json:"facadeCount"`
			StreetFacadeWidth float32 `json:"streetFacadeWidth"`
		} `json:"building"`
		Subtype string `json:"subtype"`
		Kitchen struct {
			Surface int    `json:"surface"`
			Type    string `json:"type"`
			IsSetup bool   `json:"isSetup"`
		} `json:"kitchen"`
		ShowerRoom struct {
			Count int `json:"count"`
		} `json:"showerRoom"`
		Bathroom struct {
			Count int `json:"count"`
		} `json:"bathroom"`
		Energy struct {
			HasDoubleGlazing bool   `json:"hasDoubleGlazing"`
			HeatingType      string `json:"heatingType"`
		} `json:"energy"`
		CommonEquipment struct {
			HasCaretakerOrConcierge bool `json:"hasCaretakerOrConcierge"`
			HasSecureAlarm          bool `json:"hasSecureAlarm"`
			HasLift                 bool `json:"hasLift"`
			HasDisabledAccess       bool `json:"hasDisabledAccess"`
		} `json:"commonEquipment"`
		LivingRoom struct {
			Surface int  `json:"surface"`
			Exists  bool `json:"exists"`
		} `json:"livingRoom"`
		PrivateEquipment struct {
			HasDoorPhone   bool `json:"hasDoorPhone"`
			HasArmoredDoor bool `json:"hasArmoredDoor"`
			HasVisiophone  bool `json:"hasVisiophone"`
			HasCableTV     bool `json:"hasCableTV"`
		} `json:"privateEquipment"`
		MonthlyCosts float64 `json:"monthlyCosts"`
		Bedroom      struct {
			Count int `json:"count"`
			Items []struct {
				Surface int `json:"surface"`
			} `json:"items"`
		} `json:"bedroom"`
		LivingDescription struct {
			NetHabitableSurface int `json:"netHabitableSurface"`
			Fireplace           struct {
				Count  int  `json:"count"`
				Exists bool `json:"exists"`
			} `json:"fireplace"`
		} `json:"livingDescription"`
		Outdoor struct {
			Exists  bool `json:"exists"`
			Terrace struct {
				Surface int  `json:"surface"`
				Exists  bool `json:"exists"`
			} `json:"terrace"`
		} `json:"outdoor"`
		Toilet struct {
			Count int `json:"count"`
		} `json:"toilet"`
		WellnessEquipment struct {
			HasSwimmingPool bool `json:"hasSwimmingPool"`
		} `json:"wellnessEquipment"`
		Certificates struct {
			OilTankCertificateStatus string `json:"oilTankCertificateStatus"`
		} `json:"certificates"`
		Attic struct {
			Exists bool `json:"exists"`
		} `json:"attic"`
		ConstructionPermit struct {
			TotalBuildableGroundFloorSurface   int    `json:"totalBuildableGroundFloorSurface"`
			IsBreachingUrbanPlanningRegulation bool   `json:"isBreachingUrbanPlanningRegulation"`
			HasPlotDivisionAuthorization       bool   `json:"hasPlotDivisionAuthorization"`
			FloodZoneType                      string `json:"floodZoneType"`
			IsObtained                         bool   `json:"isObtained"`
			HasPossiblePriorityPurchaseRight   bool   `json:"hasPossiblePriorityPurchaseRight"`
		} `json:"constructionPermit"`
		Location struct {
			Address struct {
				Country     string `json:"country"`
				Number      string `json:"number"`
				Province    string `json:"province"`
				Street      string `json:"street"`
				District    string `json:"district"`
				PostalCode  string `json:"postalCode"`
				Locality    string `json:"locality"`
				Box         string `json:"box"`
				Floor       int    `json:"floor"`
				Region      string `json:"region"`
				PlaceName   string `json:"placeName"`
				CountryCode string `json:"countryCode"`
				RegionCode  string `json:"regionCode"`
			} `json:"address"`
			GeoPoint struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoPoint"`
			Surroundings struct {
				PointsOfInterest []struct {
					Distance int    `json:"distance"`
					Type     string `json:"type"`
				} `json:"pointsOfInterest"`
				Type string `json:"type"`
			} `json:"surroundings"`
		} `json:"location"`
		Title             string `json:"title"`
		AlternativeTitles struct {
			En string `json:"en"`
			Nl string `json:"nl"`
		} `json:"alternativeTitles"`
		Description             string `json:"description"`
		AlternativeDescriptions struct {
			En string `json:"en"`
			Nl string `json:"nl"`
		} `json:"alternativeDescriptions"`
	} `json:"property"`
	FlagsAndStatistics struct {
		Flags struct {
			IsANewlyPublishedClassified bool `json:"isANewlyPublishedClassified"`
			IsNewlyBuilt                bool `json:"isNewlyBuilt"`
			IsAPublicSale               bool `json:"isAPublicSale"`
			IsALifeAnnuitySale          bool `json:"isALifeAnnuitySale"`
			IsAPropertyWithNewPrice     bool `json:"isAPropertyWithNewPrice"`
		} `json:"flags"`
		Statistics struct {
			Popularity    float64 `json:"popularity"`
			BookmarkCount int     `json:"bookmarkCount"`
			ViewCount     int     `json:"viewCount"`
		} `json:"statistics"`
	} `json:"flagsAndStatistics"`
	Publication struct {
		LastModificationDate string `json:"lastModificationDate"`
		VisualisationOption  string `json:"visualisationOption"`
		CreationDate         string `json:"creationDate"`
		ActivationDate       string `json:"activationDate"`
		ExpirationDate       string `json:"expirationDate"`
	} `json:"publication"`
	Customers []struct {
		ContactInfo []struct {
			Email    string `json:"email"`
			Landline string `json:"landline"`
			Usage    string `json:"usage"`
		} `json:"contactInfo"`
		ID       int `json:"id"`
		Location struct {
			Address struct {
				Country    string `json:"country"`
				District   string `json:"district"`
				Locality   string `json:"locality"`
				PostalCode string `json:"postalCode"`
				Province   string `json:"province"`
				Street     string `json:"street"`
			} `json:"address"`
			GeoPoint struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoPoint"`
		} `json:"location"`
		PublicInfo struct {
			IpiNo   string `json:"ipiNo"`
			LogoURL string `json:"logoUrl"`
			Name    string `json:"name"`
		} `json:"publicInfo"`
		Type         string `json:"type"`
		ContactHours struct {
			Landline string `json:"landline"`
			Mobile   string `json:"mobile"`
		} `json:"contactHours"`
		IsOwner bool `json:"isOwner"`
	} `json:"customers"`
	Transaction struct {
		Investor struct {
			IsInvestmentProperty bool `json:"isInvestmentProperty"`
		} `json:"investor"`
		Certificates struct {
			PrimaryEnergyConsumption struct {
				PerSqm float64 `json:"perSqm"`
			} `json:"primaryEnergyConsumption"`
			Epc struct {
				Score           string `json:"score"`
				ReportReference string `json:"reportReference"`
				LogoURL         string `json:"logoUrl"`
			} `json:"epc"`
		} `json:"certificates"`
		Subtype                string `json:"subtype"`
		Segment                string `json:"segment"`
		AvailabilityPeriodType string `json:"availabilityPeriodType"`
		AvailabilityDate       string `json:"availabilityDate"`
		Type                   string `json:"type"`
		SoldOrRented           struct {
			IsStatusDisplayed bool `json:"isStatusDisplayed"`
			IsSoldOrRented    bool `json:"isSoldOrRented"`
		} `json:"soldOrRented"`
		Rental struct {
			Residential struct {
				IsFurnished bool `json:"isFurnished"`
			} `json:"residential"`
			MonthlyRentalCosts int `json:"monthlyRentalCosts"`
			MonthlyRentalPrice int `json:"monthlyRentalPrice"`
		} `json:"rental"`
	} `json:"transaction"`
}

type AdsFromSearchQuery []struct {
	LastModificationDate string `json:"lastModificationDate"`
	ID                   int    `json:"id"`
	Media                struct {
		Pictures struct {
			BaseURL string `json:"baseUrl"`
			Items   []struct {
				Orientation string `json:"orientation"`
				RelativeURL struct {
					Large  string `json:"large"`
					Medium string `json:"medium"`
				} `json:"relativeUrl"`
			} `json:"items"`
		} `json:"pictures"`
	} `json:"media"`
	Property struct {
		Subtype  string `json:"subtype"`
		Location struct {
			Address struct {
				Country     string `json:"country"`
				Number      string `json:"number"`
				Province    string `json:"province"`
				Street      string `json:"street"`
				District    string `json:"district"`
				PostalCode  string `json:"postalCode"`
				Locality    string `json:"locality"`
				Floor       int    `json:"floor"`
				Region      string `json:"region"`
				PlaceName   string `json:"placeName"`
				CountryCode string `json:"countryCode"`
				RegionCode  string `json:"regionCode"`
			} `json:"address"`
			GeoPoint struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoPoint"`
		} `json:"location"`
		Type         string  `json:"type"`
		MonthlyCosts float64 `json:"monthlyCosts"`
		Bedroom      struct {
			Count int `json:"count"`
		} `json:"bedroom"`
		LivingDescription struct {
			NetHabitableSurface int `json:"netHabitableSurface"`
		} `json:"livingDescription"`
		Title             string `json:"title"`
		AlternativeTitles struct {
			En string `json:"en"`
			Nl string `json:"nl"`
		} `json:"alternativeTitles"`
	} `json:"property"`
	FlagsAndStatistics struct {
		Flags struct {
			IsANewlyPublishedClassified bool `json:"isANewlyPublishedClassified"`
			IsNewlyBuilt                bool `json:"isNewlyBuilt"`
			IsAPublicSale               bool `json:"isAPublicSale"`
			IsALifeAnnuitySale          bool `json:"isALifeAnnuitySale"`
			IsAPropertyWithNewPrice     bool `json:"isAPropertyWithNewPrice"`
		} `json:"flags"`
	} `json:"flagsAndStatistics"`
	Publication struct {
		LastModificationDate string `json:"lastModificationDate"`
		VisualisationOption  string `json:"visualisationOption"`
	} `json:"publication"`
	Customers []struct {
		PublicInfo struct {
			Name    string `json:"name"`
			LogoURL string `json:"logoUrl"`
		} `json:"publicInfo"`
	} `json:"customers"`
	Transaction struct {
		Certificates struct {
			Epc struct {
				Score   string `json:"score"`
				LogoURL string `json:"logoUrl"`
			} `json:"epc"`
		} `json:"certificates"`
		Subtype      string `json:"subtype"`
		Segment      string `json:"segment"`
		Type         string `json:"type"`
		SoldOrRented struct {
			IsSoldOrRented bool `json:"isSoldOrRented"`
		} `json:"soldOrRented"`
		Rental struct {
			MonthlyRentalCosts int `json:"monthlyRentalCosts"`
			MonthlyRentalPrice int `json:"monthlyRentalPrice"`
		} `json:"rental"`
	} `json:"transaction"`
}
