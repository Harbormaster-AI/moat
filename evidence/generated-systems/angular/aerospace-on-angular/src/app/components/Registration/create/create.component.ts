import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RegistrationService } from '../../../services/Registration.service';
import { Registration } from '../../../models/Registration';
import { SubBaseComponent } from '../../Registration/sub.base.component';

@Component({
    selector: 'app-create-registration',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRegistrationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Registration';

    registrationForm: FormGroup;
    registration: Registration;

    constructor( http: HttpClient,
        private registrationService: RegistrationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRegistration(tailNumber, registryCountry, Aircraft): void {
        this.registrationService
        .addRegistration(tailNumber, registryCountry, Aircraft)
            .subscribe(() => {
                this.router.navigate(['/indexRegistration']);
            });
    }

    ngOnInit(): void {
    }
}