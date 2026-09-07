
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ObligationService } from '../../../services/Obligation.service';
import { Obligation } from '../../../models/Obligation';

@Component({
    selector: 'app-index-obligation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexObligationComponent implements OnInit {

    obligations: Obligation[] = [];

    constructor(
        private router: Router,
        private service: ObligationService
) {}

    ngOnInit(): void {
        this.getObligations();
}

    getObligations(): void {
        this.service.getObligations().subscribe((res) => {
        this.obligations = res;
    });
}

    deleteObligation(id: any): void {
        this.service.deleteObligation(id)
            .subscribe(() => {
                this.getObligations();
            });
    }
}