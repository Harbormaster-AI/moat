import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HealthSystemService } from '../../../services/HealthSystem.service';
import { HealthSystem } from '../../../models/HealthSystem';
import { SubBaseComponent } from '../../HealthSystem/sub.base.component';

@Component({
    selector: 'app-create-healthSystem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateHealthSystemComponent extends SubBaseComponent implements OnInit {

    title = 'Add HealthSystem';

    healthSystemForm: FormGroup;
    healthSystem: HealthSystem;

    constructor( http: HttpClient,
        private healthSystemService: HealthSystemService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.healthSystemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      headquartersCountry: ['', Validators.required],
      website: ['', Validators.required],
      Facilities: ['', ],
      Suppliers: ['', ]
        });
    }

    
    addHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers): void {
        this.healthSystemService
        .addHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers)
            .subscribe(() => {
                this.router.navigate(['/indexHealthSystem']);
            });
    }

    ngOnInit(): void {
    }
}