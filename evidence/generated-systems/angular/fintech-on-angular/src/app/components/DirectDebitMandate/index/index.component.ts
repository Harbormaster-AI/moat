
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';
import { DirectDebitMandate } from '../../../models/DirectDebitMandate';

@Component({
    selector: 'app-index-directDebitMandate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDirectDebitMandateComponent implements OnInit {

    directDebitMandates: DirectDebitMandate[] = [];

    constructor(
        private router: Router,
        private service: DirectDebitMandateService
) {}

    ngOnInit(): void {
        this.getDirectDebitMandates();
}

    getDirectDebitMandates(): void {
        this.service.getDirectDebitMandates().subscribe((res) => {
        this.directDebitMandates = res;
    });
}

    deleteDirectDebitMandate(id: any): void {
        this.service.deleteDirectDebitMandate(id)
            .subscribe(() => {
                this.getDirectDebitMandates();
            });
    }
}