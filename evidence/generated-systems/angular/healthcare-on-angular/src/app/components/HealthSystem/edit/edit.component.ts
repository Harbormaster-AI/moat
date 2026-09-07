import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { HealthSystemService } from '../../../services/HealthSystem.service';
import { SubBaseComponent } from '../../HealthSystem/sub.base.component';


@Component({
    selector: 'app-edit-healthSystem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditHealthSystemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit HealthSystem';

    healthSystemForm: FormGroup;
    healthSystem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: HealthSystemService,
        private fb: FormBuilder
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

    
    updateHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers): void {
        this.route.params.subscribe((params) => {

                        this.service.updateHealthSystem(name, legalName, headquartersCountry, website, Facilities, Suppliers, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexHealthSystem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getHealthSystem(params['id']).subscribe(res => {
                this.healthSystem = res;
            });
        });
    }
}