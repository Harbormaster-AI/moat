
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ChargebackService } from '../../../services/Chargeback.service';
import { Chargeback } from '../../../models/Chargeback';

@Component({
    selector: 'app-index-chargeback',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexChargebackComponent implements OnInit {

    chargebacks: Chargeback[] = [];

    constructor(
        private router: Router,
        private service: ChargebackService
) {}

    ngOnInit(): void {
        this.getChargebacks();
}

    getChargebacks(): void {
        this.service.getChargebacks().subscribe((res) => {
        this.chargebacks = res;
    });
}

    deleteChargeback(id: any): void {
        this.service.deleteChargeback(id)
            .subscribe(() => {
                this.getChargebacks();
            });
    }
}