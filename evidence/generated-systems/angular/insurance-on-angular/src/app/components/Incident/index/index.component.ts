
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { IncidentService } from '../../../services/Incident.service';
import { Incident } from '../../../models/Incident';

@Component({
    selector: 'app-index-incident',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexIncidentComponent implements OnInit {

    incidents: Incident[] = [];

    constructor(
        private router: Router,
        private service: IncidentService
) {}

    ngOnInit(): void {
        this.getIncidents();
}

    getIncidents(): void {
        this.service.getIncidents().subscribe((res) => {
        this.incidents = res;
    });
}

    deleteIncident(id: any): void {
        this.service.deleteIncident(id)
            .subscribe(() => {
                this.getIncidents();
            });
    }
}