import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { UoMConversionService } from '../../../services/UoMConversion.service';
import { SubBaseComponent } from '../../UoMConversion/sub.base.component';


@Component({
    selector: 'app-edit-uoMConversion',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditUoMConversionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit UoMConversion';

    uoMConversionForm: FormGroup;
    uoMConversion: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: UoMConversionService,
        private fb: FormBuilder
) {
        super(http);
        this.uoMConversionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  factor: ['', Validators.required],
      precision: ['', Validators.required],
      Sku: ['', ],
      FromUnit: ['', ],
      ToUnit: ['', ]
        });
    }

    
    updateUoMConversion(factor, precision, Sku, FromUnit, ToUnit): void {
        this.route.params.subscribe((params) => {

                        this.service.updateUoMConversion(factor, precision, Sku, FromUnit, ToUnit, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexUoMConversion']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getUoMConversion(params['id']).subscribe(res => {
                this.uoMConversion = res;
            });
        });
    }
}