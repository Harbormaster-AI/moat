
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { KPIService } from '../../../services/KPI.service';
import { KPI } from '../../../models/KPI';

@Component({
    selector: 'app-index-kPI',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexKPIComponent implements OnInit {

    kPIs: KPI[] = [];

    constructor(
        private router: Router,
        private service: KPIService
) {}

    ngOnInit(): void {
        this.getKPIs();
}

    getKPIs(): void {
        this.service.getKPIs().subscribe((res) => {
        this.kPIs = res;
    });
}

    deleteKPI(id: any): void {
        this.service.deleteKPI(id)
            .subscribe(() => {
                this.getKPIs();
            });
    }
}