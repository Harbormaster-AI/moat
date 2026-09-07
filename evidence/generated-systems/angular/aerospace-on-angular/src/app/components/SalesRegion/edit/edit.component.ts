import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SalesRegionService } from '../../../services/SalesRegion.service';
import { SubBaseComponent } from '../../SalesRegion/sub.base.component';


@Component({
    selector: 'app-edit-salesRegion',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSalesRegionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SalesRegion';

    salesRegionForm: FormGroup;
    salesRegion: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SalesRegionService,
        private fb: FormBuilder
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

    
    updateSalesRegion(name, regionCode, Operators, SalesCampaigns): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSalesRegion(name, regionCode, Operators, SalesCampaigns, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSalesRegion']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSalesRegion(params['id']).subscribe(res => {
                this.salesRegion = res;
            });
        });
    }
}