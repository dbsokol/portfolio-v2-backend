INSERT INTO experiences (company, role, mission, start_date) VALUES
('Tech Corp', 'Software Engineer', 'Develop scalable backend systems', '2020-01-01'),
('Innovate Labs', 'Data Scientist', 'Analyze and optimize ML pipelines', '2018-05-01');

INSERT INTO responsibilities (experience_id, description) VALUES
(1, 'Built REST APIs in Go'),
(2, 'Implemented predictive analytics models');

INSERT INTO skills (name) VALUES
('Golang'),
('Machine Learning');
