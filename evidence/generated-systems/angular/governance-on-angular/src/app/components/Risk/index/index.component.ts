
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RiskService } from '../../../services/Risk.service';
import { Risk } from '../../../models/Risk';

@Component({
    selector: 'app-index-risk',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRiskComponent implements OnInit {

    risks: Risk[] = [];

    constructor(
        private router: Router,
        private service: RiskService
) {}

    ngOnInit(): void {
        this.getRisks();
}

    getRisks(): void {
        this.service.getRisks().subscribe((res) => {
        this.risks = res;
    });
}

    deleteRisk(id: any): void {
        this.service.deleteRisk(id)
            .subscribe(() => {
                this.getRisks();
            });
    }
}