import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SupplierService } from '../../../services/Supplier.service';
import { Supplier } from '../../../models/Supplier';
import { SubBaseComponent } from '../../Supplier/sub.base.component';

@Component({
    selector: 'app-create-supplier',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSupplierComponent extends SubBaseComponent implements OnInit {

    title = 'Add Supplier';

    supplierForm: FormGroup;
    supplier: Supplier;

    constructor( http: HttpClient,
        private supplierService: SupplierService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.supplierForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Manufacturers: ['', ],
      Components: ['', ],
      EngineTypes: ['', ],
      AvionicsSuites: ['', ],
      Apus: ['', ],
      LandingGears: ['', ],
      SupplierType: ['', ],
      ApprovalStatus: ['', ]
        });
    }

    
    addSupplier(name, Manufacturers, Components, EngineTypes, AvionicsSuites, Apus, LandingGears, SupplierType, ApprovalStatus): void {
        this.supplierService
        .addSupplier(name, Manufacturers, Components, EngineTypes, AvionicsSuites, Apus, LandingGears, SupplierType, ApprovalStatus)
            .subscribe(() => {
                this.router.navigate(['/indexSupplier']);
            });
    }

    ngOnInit(): void {
    }
}