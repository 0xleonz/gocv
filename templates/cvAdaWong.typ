#import "@preview/simple-technical-resume:0.1.0": *

// Put your personal information here
#let name = "Ada Wong"
#let phone = "+1 (555) 555-1998"
#let email = "awong@binom.com"
#let github = "ada-wong-spy"
#let linkedin = "ada-wong"
#let personal-site = "binom-espionage.com"

// Since the following arguments are within the `with` block,
// you can remove/comment any argument to fallback to the preset value and/or
// remove it. 
#show: resume.with(
  top-margin: 0.45in,
  font: "New Computer Modern",
  personal-info-font-size: 9.2pt,
  author-position: center,
  personal-info-position: center,
  author-name: name,
  phone: phone,
  email: email,
  website: personal-site,
  linkedin-user-id: linkedin,
  github-username: github
)

// Use custom-title function instead of first-level headings to customize the
// size between two sections by specifying the `spacingBetween` argument.
// https://typst.app/docs/reference/layout/length/

#custom-title("Education")[
  #education-heading(
    "Classified Intelligence Academy",           // institution
    "Unknown Location",                         // location
    "Master of Espionage Arts",                 // degree
    "Infiltration & Exfiltration",              // major
    datetime(year: 1998, month: 9, day: 1),     // start-date
    datetime(year: 2000, month: 6, day: 1)      // end-date
  )[
    - Specialized in covert operations, disguise, and advanced combat techniques.
  ]
  #education-heading(
    "University of B.O.W. Studies",             // institution
    "Unknown Location",                         // location
    "Doctorate (Honorary)",                     // degree
    "Bio-Organic Weapons Analysis",             // major
    datetime(year: 2001, month: 9, day: 1),     // start-date
    datetime(year: 2004, month: 6, day: 1)      // end-date
  )[
    - Developed expertise in identifying, analyzing, and neutralizing bio-organic threats.
  ]
]

#custom-title("Experience")[  
  #work-heading(
    "Independent Operative / Corporate Saboteur",// title
    "Various Organizations (e.g., The Family, Tricell)", // company
    "Global",                                   // location
    datetime(year: 2004, month: 6, day: 1),     // start-date
    "Present"                                   // end-date
  )[
    - Successfully infiltrated numerous high-security facilities to acquire or neutralize dangerous biological agents and research.
    - Masterfully employed deception and manipulation to achieve mission objectives with minimal direct engagement.
    - Provided crucial intel and support to key figures during global B.O.W. crises.
    - Adept at disappearing without a trace after mission completion.
  ]
  #work-heading(
    "Field Agent",                              // title
    "Unknown Espionage Agency",                 // company
    "Global",                                   // location
    datetime(year: 1998, month: 6, day: 1),     // start-date
    datetime(year: 2004, month: 6, day: 1)      // end-date
  )[
    - Conducted numerous covert assignments, including data retrieval, assassination, and sabotage.
    - Demonstrated exceptional adaptability in high-pressure, combat-intensive scenarios.
  ]
]

#custom-title("Projects")[
  #project-heading(
    "Operation NE-a (Resident Evil 2)",          // name
    // "Next.js, TailwindCSS, Postgres",          // stack
    // "schrutefarms.com"                         // project_url
  )[
    - Infiltrated Raccoon City during its outbreak to retrieve a G-Virus sample for her employer.
    - Navigated treacherous environments and evaded hostile forces, including the Tyrant T-00.
  ]
  #project-heading(
    "Operation 'Las Plagas' Neutralization",     // name
    // "Next.js, TailwindCSS, Postgres",          // stack
    // "schrutefarms.com"                         // project_url
  )[
    - Infiltrated a Spanish village controlled by a cult to recover a sample of the Las Plagas parasite.
    - Provided critical assistance to Leon S. Kennedy, ensuring mission success.
  ]
  // More projects ...
]

// Use `skills` function to create list with custom rules surrounding indentation and alignment.
// It is specifically for lists directly inside the custom-title section.
#custom-title("Skills")[
  #skills()[
    - *Professional Skills:* Espionage, Infiltration, Exfiltration, Disguise, Combat (Hand-to-Hand & Firearms), Interrogation, Sabotage
    - *Technical Skills:* Hacking, Lockpicking, Gadget Use (grappling hook, crossbow), Driving (various vehicles)
    - *Specialized Talents:* Master of deception, fluent in multiple languages, exceptional agility and reflexes, survival under extreme conditions
  ]
]
