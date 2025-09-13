#import "@preview/simple-technical-resume:0.1.0": *

// Put your personal information here
#let name = "Mario Mario"
#let phone = "+1 (555) 555-1985"
#let email = "mmario@mushroomkingdom.com"
#let github = "super-mario"
#let linkedin = "mario-mario"
#let personal-site = "toad-ally-awesome.com"

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
    "Mushroom Kingdom Plumbing Academy",     // institution
    "Toad Town, MK",                         // location
    "Associate of Applied Science",          // degree
    "Piping & Drainage",                     // major
    datetime(year: 1981, month: 9, day: 1),  // start-date
    datetime(year: 1983, month: 6, day: 1)   // end-date
  )[
    - Graduated at the top of the class in "Pipe Navigation and Warp Zone Theory"
  ]
  // More educational qualifications ... 
]

#custom-title("Experience")[  
  #work-heading(
    "Freelance Hero & Royal Rescuer",       // title
    "Mushroom Kingdom",                     // company
    "Various Locations",                    // location
    datetime(year: 1985, month: 9, day: 13),// start-date
    "Present"                               // end-date
  )[
    - Consistently rescued Princess Peach and the Mushroom Kingdom from Bowser's invasions, demonstrating exceptional bravery and strategic thinking
    - Successfully navigated treacherous environments, including lava pits, haunted mansions, and sky fortresses
    - Utilized a wide range of power-ups to overcome obstacles and defeat enemies, including Fire Flowers, Super Stars, and Tanooki Suits
    - Collaborated with allies like Luigi, Toad, and Yoshi to achieve objectives and maintain peace
  ]
  #work-heading(
    "Owner & Operator",                     // title
    "Mario Bros. Plumbing",                 // company
    "Brooklyn, NY",                         // location
    datetime(year: 1983, month: 6, day: 1), // start-date
    datetime(year: 1985, month: 9, day: 12)// end-date
  )[
    - Founded and operated a successful plumbing business with my brother, Luigi
    - Specialized in high-stakes, subterranean pipe work, often dealing with unique biological and environmental hazards
    - Maintained a perfect customer satisfaction record
  ]
]

#custom-title("Projects")[
  #project-heading(
    "Mario Kart Circuit Design",             // name
    // "Next.js, TailwindCSS, Postgres",       // stack
    // "schrutefarms.com"                       // project_url
  )[
    - Conceived and designed numerous high-speed race tracks, including Rainbow Road and Wario's Gold Mine
    - Developed and tested innovative vehicle and kart modifications to enhance performance
  ]
  #project-heading(
    "Super Smash Bros. Tournament Series",   // name
    // "Next.js, TailwindCSS, Postgres",       // stack
    // "schrutefarms.com"                       // project_url
  )[
    - Participated in and won multiple inter-dimensional combat tournaments, showcasing a versatile fighting style
    - Served as a foundational character, inspiring countless others to join the fray
  ]
  // More projects ...
]

// Use `skills` function to create list with custom rules surrounding indentation and alignment.
// It is specifically for lists directly inside the custom-title section.
#custom-title("Skills")[
  #skills()[
    - *Professional Skills:* Heroism, Plunging, Brick-Breaking, Goomba Stomping, Platforming, Strategic Planning
    - *Power-Ups:* Fire Flower, Super Star, Super Mushroom, Tanooki Suit, Cape Feather, Frog Suit
    - *Specialized Talents:* High Jump, Wall Jump, Triple Jump, Kart Racing, Golf, Tennis, Partying
  ]
]
