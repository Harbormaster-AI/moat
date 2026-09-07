
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LegalHoldService } from '../../../services/LegalHold.service';
import { LegalHold } from '../../../models/LegalHold';

@Component({
    selector: 'app-index-legalHold',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLegalHoldComponent implements OnInit {

    legalHolds: LegalHold[] = [];

    constructor(
        private router: Router,
        private service: LegalHoldService
) {}

    ngOnInit(): void {
        this.getLegalHolds();
}

    getLegalHolds(): void {
        this.service.getLegalHolds().subscribe((res) => {
        this.legalHolds = res;
    });
}

    deleteLegalHold(id: any): void {
        this.service.deleteLegalHold(id)
            .subscribe(() => {
                this.getLegalHolds();
            });
    }
}