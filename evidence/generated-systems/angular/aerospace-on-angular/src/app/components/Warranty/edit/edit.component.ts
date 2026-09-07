import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WarrantyService } from '../../../services/Warranty.service';
import { SubBaseComponent } from '../../Warranty/sub.base.component';


@Component({
    selector: 'app-edit-warranty',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWarrantyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Warranty';

    warrantyForm: FormGroup;
    warranty: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WarrantyService,
        private fb: FormBuilder
) {
        super(http);
        this.warrantyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  coverageMonths: ['', Validators.required],
      Aircraft: ['', ],
      WarrantyType: ['', ]
        });
    }

    
    updateWarranty(coverageMonths, Aircraft, WarrantyType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWarranty(coverageMonths, Aircraft, WarrantyType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWarranty']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWarranty(params['id']).subscribe(res => {
                this.warranty = res;
            });
        });
    }
}