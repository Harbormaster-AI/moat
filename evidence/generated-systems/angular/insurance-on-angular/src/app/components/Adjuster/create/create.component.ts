import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdjusterService } from '../../../services/Adjuster.service';
import { Adjuster } from '../../../models/Adjuster';
import { SubBaseComponent } from '../../Adjuster/sub.base.component';

@Component({
    selector: 'app-create-adjuster',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAdjusterComponent extends SubBaseComponent implements OnInit {

    title = 'Add Adjuster';

    adjusterForm: FormGroup;
    adjuster: Adjuster;

    constructor( http: HttpClient,
        private adjusterService: AdjusterService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.adjusterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      licenseNumber: ['', Validators.required],
      Claims: ['', ],
      ServiceProviders: ['', ],
      AdjusterType: ['', ]
        });
    }

    
    addAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType): void {
        this.adjusterService
        .addAdjuster(firstName, lastName, licenseNumber, Claims, ServiceProviders, AdjusterType)
            .subscribe(() => {
                this.router.navigate(['/indexAdjuster']);
            });
    }

    ngOnInit(): void {
    }
}