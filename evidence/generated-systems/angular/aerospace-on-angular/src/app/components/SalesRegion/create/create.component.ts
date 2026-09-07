import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SalesRegionService } from '../../../services/SalesRegion.service';
import { SalesRegion } from '../../../models/SalesRegion';
import { SubBaseComponent } from '../../SalesRegion/sub.base.component';

@Component({
    selector: 'app-create-salesRegion',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSalesRegionComponent extends SubBaseComponent implements OnInit {

    title = 'Add SalesRegion';

    salesRegionForm: FormGroup;
    salesRegion: SalesRegion;

    constructor( http: HttpClient,
        private salesRegionService: SalesRegionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.salesRegionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      regionCode: ['', Validators.required],
      Operators: ['', ],
      SalesCampaigns: ['', ]
        });
    }

    
    addSalesRegion(name, regionCode, Operators, SalesCampaigns): void {
        this.salesRegionService
        .addSalesRegion(name, regionCode, Operators, SalesCampaigns)
            .subscribe(() => {
                this.router.navigate(['/indexSalesRegion']);
            });
    }

    ngOnInit(): void {
    }
}