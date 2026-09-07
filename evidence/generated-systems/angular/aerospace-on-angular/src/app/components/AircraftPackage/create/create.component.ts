import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftPackageService } from '../../../services/AircraftPackage.service';
import { AircraftPackage } from '../../../models/AircraftPackage';
import { SubBaseComponent } from '../../AircraftPackage/sub.base.component';

@Component({
    selector: 'app-create-aircraftPackage',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftPackageComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftPackage';

    aircraftPackageForm: FormGroup;
    aircraftPackage: AircraftPackage;

    constructor( http: HttpClient,
        private aircraftPackageService: AircraftPackageService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.aircraftPackageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Options: ['', ],
      Variants: ['', ],
      PackageType: ['', ]
        });
    }

    
    addAircraftPackage(name, Options, Variants, PackageType): void {
        this.aircraftPackageService
        .addAircraftPackage(name, Options, Variants, PackageType)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftPackage']);
            });
    }

    ngOnInit(): void {
    }
}