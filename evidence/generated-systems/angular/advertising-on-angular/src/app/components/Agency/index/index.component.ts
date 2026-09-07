
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgencyService } from '../../../services/Agency.service';
import { Agency } from '../../../models/Agency';

@Component({
    selector: 'app-index-agency',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAgencyComponent implements OnInit {

    agencys: Agency[] = [];

    constructor(
        private router: Router,
        private service: AgencyService
) {}

    ngOnInit(): void {
        this.getAgencys();
}

    getAgencys(): void {
        this.service.getAgencys().subscribe((res) => {
        this.agencys = res;
    });
}

    deleteAgency(id: any): void {
        this.service.deleteAgency(id)
            .subscribe(() => {
                this.getAgencys();
            });
    }
}