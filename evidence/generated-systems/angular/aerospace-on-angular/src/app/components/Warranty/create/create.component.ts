import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WarrantyService } from '../../../services/Warranty.service';
import { Warranty } from '../../../models/Warranty';
import { SubBaseComponent } from '../../Warranty/sub.base.component';

@Component({
    selector: 'app-create-warranty',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWarrantyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Warranty';

    warrantyForm: FormGroup;
    warranty: Warranty;

    constructor( http: HttpClient,
        private warrantyService: WarrantyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addWarranty(coverageMonths, Aircraft, WarrantyType): void {
        this.warrantyService
        .addWarranty(coverageMonths, Aircraft, WarrantyType)
            .subscribe(() => {
                this.router.navigate(['/indexWarranty']);
            });
    }

    ngOnInit(): void {
    }
}