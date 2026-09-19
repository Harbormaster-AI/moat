class WorkShift < ApplicationRecord
  enum DayOfWeek: [:Monday, :Tuesday, :Wednesday, :Thursday, :Friday, :Saturday, :Sunday]


  composed_of :localTime,
    class_name: "LocalTime",
    mapping: [
      ${$mapping}, 
      %w[localTime_minute minute]
    ]

  composed_of :localTime,
    class_name: "LocalTime",
    mapping: [
      ${$mapping}, 
      %w[localTime_minute minute]
    ]

  has_many :WorkSchedule, class_name: 'WorkSchedule'

end
