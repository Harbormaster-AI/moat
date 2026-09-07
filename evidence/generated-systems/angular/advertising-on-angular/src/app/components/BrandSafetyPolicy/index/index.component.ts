
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';
import { BrandSafetyPolicy } from '../../../models/BrandSafetyPolicy';

@Component({
    selector: 'app-index-brandSafetyPolicy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBrandSafetyPolicyComponent implements OnInit {

    brandSafetyPolicys: BrandSafetyPolicy[] = [];

    constructor(
        private router: Router,
        private service: BrandSafetyPolicyService
) {}

    ngOnInit(): void {
        this.getBrandSafetyPolicys();
}

    getBrandSafetyPolicys(): void {
        this.service.getBrandSafetyPolicys().subscribe((res) => {
        this.brandSafetyPolicys = res;
    });
}

    deleteBrandSafetyPolicy(id: any): void {
        this.service.deleteBrandSafetyPolicy(id)
            .subscribe(() => {
                this.getBrandSafetyPolicys();
            });
    }
}