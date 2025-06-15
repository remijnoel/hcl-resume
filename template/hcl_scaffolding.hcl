basics {
  name    = "" # "Firstname Lastname"
  label   = "" # "Name of your role, title, etc."
  email   = ""
  phone   = ""
  url     = ""
  summary = "" # Your professional summary
  image   = ""
  location {
    address      = ""
    postal_code  = ""
    city         = ""
    country_code = ""
    region       = ""
  }
}


work "a_name_for_this_role" {
  company    = ""
  position   = "" # Name of the position
#   website    = "" # Optional
  start_date = ""
  end_date   = ""
  summary    = "" # Brief summary of your role
  highlights = [
    "" # Key achievements or responsibilities - the bullet points for the role
  ]
  location = ""
}


education "a_name_for_this_education" {
  institution = "" # Name of the institution
  area        = "" # The area of study e.g. Computer Science
  study_type  = "" # e.g. Bachelor, Master, PhD
  start_date  = ""
  end_date    = ""
#   courses     = [""] # List of relevant courses
}


skills "a_name_for_this_skill" {
  name     = "" # The category of skills, e.g. "Cloud Platforms"
  level    = "expert"
  keywords = ["aws", "azure", "gcp", "terraform", "docker"]
}


certificate "cloud architect cert" {
  issuer = "" # Name of the certifying body
  date   = ""
  url    = "" # URL to the certificate or more info
}
