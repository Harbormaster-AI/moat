
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';
import { VerifiedAddress } from '../../../models/VerifiedAddress';

@Component({
    selector: 'app-index-verifiedAddress',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexVerifiedAddressComponent implements OnInit {

    verifiedAddresss: VerifiedAddress[] = [];

    constructor(
        private router: Router,
        private service: VerifiedAddressService
) {}

    ngOnInit(): void {
        this.getVerifiedAddresss();
}

    getVerifiedAddresss(): void {
        this.service.getVerifiedAddresss().subscribe((res) => {
        this.verifiedAddresss = res;
    });
}

    deleteVerifiedAddress(id: any): void {
        this.service.deleteVerifiedAddress(id)
            .subscribe(() => {
                this.getVerifiedAddresss();
            });
    }
}