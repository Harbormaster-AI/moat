import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RegistrationService } from '../../../services/Registration.service';
import { SubBaseComponent } from '../../Registration/sub.base.component';


@Component({
    selector: 'app-edit-registration',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRegistrationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Registration';

    registrationForm: FormGroup;
    registration: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RegistrationService,
        private fb: FormBuilder
) {
        super(http);
        this.registrationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  tailNumber: ['', Validators.required],
      registryCountry: ['', Validators.required],
      Aircraft: ['', ]
        });
    }

    
    updateRegistration(tailNumber, registryCountry, Aircraft): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRegistration(tailNumber, registryCountry, Aircraft, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRegistration']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRegistration(params['id']).subscribe(res => {
                this.registration = res;
            });
        });
    }
}