
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LeadService } from '../../../services/Lead.service';
import { Lead } from '../../../models/Lead';

@Component({
    selector: 'app-index-lead',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLeadComponent implements OnInit {

    leads: Lead[] = [];

    constructor(
        private router: Router,
        private service: LeadService
) {}

    ngOnInit(): void {
        this.getLeads();
}

    getLeads(): void {
        this.service.getLeads().subscribe((res) => {
        this.leads = res;
    });
}

    deleteLead(id: any): void {
        this.service.deleteLead(id)
            .subscribe(() => {
                this.getLeads();
            });
    }
}