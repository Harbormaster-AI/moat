import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UoMConversionService } from '../../../services/UoMConversion.service';
import { UoMConversion } from '../../../models/UoMConversion';
import { SubBaseComponent } from '../../UoMConversion/sub.base.component';

@Component({
    selector: 'app-create-uoMConversion',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateUoMConversionComponent extends SubBaseComponent implements OnInit {

    title = 'Add UoMConversion';

    uoMConversionForm: FormGroup;
    uoMConversion: UoMConversion;

    constructor( http: HttpClient,
        private uoMConversionService: UoMConversionService,
        private fb: FormBuilder,
        private router: Router
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

    
    addUoMConversion(factor, precision, Sku, FromUnit, ToUnit): void {
        this.uoMConversionService
        .addUoMConversion(factor, precision, Sku, FromUnit, ToUnit)
            .subscribe(() => {
                this.router.navigate(['/indexUoMConversion']);
            });
    }

    ngOnInit(): void {
    }
}