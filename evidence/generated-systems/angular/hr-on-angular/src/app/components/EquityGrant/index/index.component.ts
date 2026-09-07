
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EquityGrantService } from '../../../services/EquityGrant.service';
import { EquityGrant } from '../../../models/EquityGrant';

@Component({
    selector: 'app-index-equityGrant',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEquityGrantComponent implements OnInit {

    equityGrants: EquityGrant[] = [];

    constructor(
        private router: Router,
        private service: EquityGrantService
) {}

    ngOnInit(): void {
        this.getEquityGrants();
}

    getEquityGrants(): void {
        this.service.getEquityGrants().subscribe((res) => {
        this.equityGrants = res;
    });
}

    deleteEquityGrant(id: any): void {
        this.service.deleteEquityGrant(id)
            .subscribe(() => {
                this.getEquityGrants();
            });
    }
}