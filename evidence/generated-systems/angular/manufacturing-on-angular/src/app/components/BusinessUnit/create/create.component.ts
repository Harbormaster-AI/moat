import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BusinessUnitService } from '../../../services/BusinessUnit.service';
import { BusinessUnit } from '../../../models/BusinessUnit';
import { SubBaseComponent } from '../../BusinessUnit/sub.base.component';

@Component({
    selector: 'app-create-businessUnit',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBusinessUnitComponent extends SubBaseComponent implements OnInit {

    title = 'Add BusinessUnit';

    businessUnitForm: FormGroup;
    businessUnit: BusinessUnit;

    constructor( http: HttpClient,
        private businessUnitService: BusinessUnitService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.businessUnitForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      code: ['', Validators.required],
      Enterprise: ['', ],
      Items: ['', ],
      Plants: ['', ],
      Category: ['', ]
        });
    }

    
    addBusinessUnit(name, code, Enterprise, Items, Plants, Category): void {
        this.businessUnitService
        .addBusinessUnit(name, code, Enterprise, Items, Plants, Category)
            .subscribe(() => {
                this.router.navigate(['/indexBusinessUnit']);
            });
    }

    ngOnInit(): void {
    }
}