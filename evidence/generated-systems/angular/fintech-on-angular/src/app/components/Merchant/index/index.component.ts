
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MerchantService } from '../../../services/Merchant.service';
import { Merchant } from '../../../models/Merchant';

@Component({
    selector: 'app-index-merchant',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMerchantComponent implements OnInit {

    merchants: Merchant[] = [];

    constructor(
        private router: Router,
        private service: MerchantService
) {}

    ngOnInit(): void {
        this.getMerchants();
}

    getMerchants(): void {
        this.service.getMerchants().subscribe((res) => {
        this.merchants = res;
    });
}

    deleteMerchant(id: any): void {
        this.service.deleteMerchant(id)
            .subscribe(() => {
                this.getMerchants();
            });
    }
}