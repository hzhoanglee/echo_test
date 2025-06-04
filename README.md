[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/1HPWjYsx)

# ECHO: Smart Device Management System

## Team Members
- 2440045 - Dang Thai Son
- 2440040 - Nguyen Huy Hoang
- 2440041 - Bui Tuan Minh
- 2440046 - Phi Doan Minh Luong
- 2440051 - Le Minh Hoang

## Overview

ECHO (Smart-home Device Operations and Control Kit) is a comprehensive smart device management system with the special ability to operate when internet connection is lost. This project focuses on developing a reliable, secure, and energy-efficient smart home solution with intelligent decision-making capabilities based on the PEAS model (Performance, Environment, Actuators, Sensors).

## PEAS Model

### Performance
- Response latency under 2000ms
- Ability to operate in local networks
- Continuous operation during connection loss
- Energy consumption optimization

### Environment
- Indoor operating range (10m when network is lost)
- Monitoring of rooms and areas within the home
- Tracking temperature and humidity conditions
- Monitoring door status

### Actuators
- ESP32 connected to fans
- ESP32 connected to lights
- ESP32 connected to TV
- Other smart switches

### Sensors
- Temperature sensors
- Door status sensors
- Motion sensors
- Other monitoring devices

## Implementation Plan

### Phase 1: Basic Development
- Setup CCU and database
- Design Flutter app UI/UX
- Develop API endpoints
- Basic ESP32 device programming
- WiFi connection testing

### Phase 2: Failover Mechanism
- Implementation of ESP-NOW
- Failover feature development
- Testing transitions during connection loss
- Optimization of operating range

### Phase 3: Intelligent Processing
- Building knowledge base
- Implementation of decision logic
- Development of prediction algorithms
- Development of smart data display widgets
- Testing intelligent decision-making

### Phase 4: Security and Refinement
- Data encryption implementation
- Authentication system setup
- Flutter app security testing
- User interface optimization

## Future Prospects

- Integration of artificial intelligence to improve prediction capabilities
- Expansion of the system with voice control
- Development of additional custom modules
- Integration with other smart home ecosystems

---

*ECHO Smart Home Project - April 13, 2025*