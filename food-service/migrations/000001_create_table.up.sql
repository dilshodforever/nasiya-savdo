CREATE TYPE incident_status AS ENUM ('reported', 'responding', 'resolved');

CREATE TABLE emergency_incidents (
    incident_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    incident_type VARCHAR(50) NOT NULL,
    location VARCHAR(50),
    description TEXT NOT NULL,
    status incident_status NOT NULL,
    reported_at TIMESTAMP DEFAULT now()
);

CREATE TYPE resource_status AS ENUM ('available', 'dispatched', 'out_of_service');

CREATE TABLE emergency_resources (
    resource_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    resource_type VARCHAR(50) NOT NULL,
    current_location VARCHAR(50),
    status resource_status NOT NULL
);

CREATE TABLE resource_dispatches (
    dispatch_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    incident_id UUID,
    resource_id UUID,
    dispatched_at TIMESTAMP NOT NULL,
    arrived_at TIMESTAMP
);

CREATE TABLE emergency_alerts (
    alert_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alert_type VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    affected_area VARCHAR(50),
    issued_at TIMESTAMP DEFAULT now()
);

CREATE TABLE evacuation_routes (
    route_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    start_point VARCHAR(50) NOT NULL,
    end_point varchar(50) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE emergency_drills (
    drill_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    drill_type VARCHAR(50) NOT NULL,
    location varchar(30) NOT NULL,
    scheduled_at TIMESTAMP  NOT NULL,
    completed_at TIMESTAMP
);

CREATE TABLE emergency_facilities (
    facility_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    facility_type VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    capacity INTEGER NOT NULL
);

INSERT INTO emergency_facilities (facility_type, name, address, capacity) VALUES
('Hospital', 'General Hospital', '123 Main St, Springfield', 200),
('Fire Station', 'Downtown Fire Station', '456 Elm St, Springfield', 50),
('Police Station', 'Central Police Station', '789 Oak St, Springfield', 100),
('Shelter', 'Community Shelter', '101 Maple St, Springfield', 150),
('Clinic', 'Downtown Clinic', '202 Pine St, Springfield', 80);


INSERT INTO evacuation_routes (start_point, end_point, description, affected_area) VALUES
('A', 'B', 'Route from A to B through the main road', 'Urban area'),
('C', 'D', 'Route from C to D along the river', 'Riverside'),
('E', 'F', 'Route from E to F via the highway', 'Highway'),
('G', 'H', 'Scenic route from G to H through the mountains', 'Mountainous region'),
('I', 'J', 'Short route from I to J through the city center', 'Urban area');
