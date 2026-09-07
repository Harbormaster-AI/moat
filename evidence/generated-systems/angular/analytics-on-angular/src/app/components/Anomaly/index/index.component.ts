
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AnomalyService } from '../../../services/Anomaly.service';
import { Anomaly } from '../../../models/Anomaly';

@Component({
    selector: 'app-index-anomaly',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAnomalyComponent implements OnInit {

    anomalys: Anomaly[] = [];

    constructor(
        private router: Router,
        private service: AnomalyService
) {}

    ngOnInit(): void {
        this.getAnomalys();
}

    getAnomalys(): void {
        this.service.getAnomalys().subscribe((res) => {
        this.anomalys = res;
    });
}

    deleteAnomaly(id: any): void {
        this.service.deleteAnomaly(id)
            .subscribe(() => {
                this.getAnomalys();
            });
    }
}