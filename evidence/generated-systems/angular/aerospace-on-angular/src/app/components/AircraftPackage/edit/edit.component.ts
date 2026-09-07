import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftPackageService } from '../../../services/AircraftPackage.service';
import { SubBaseComponent } from '../../AircraftPackage/sub.base.component';


@Component({
    selector: 'app-edit-aircraftPackage',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftPackageComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftPackage';

    aircraftPackageForm: FormGroup;
    aircraftPackage: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftPackageService,
        private fb: FormBuilder
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

    
    updateAircraftPackage(name, Options, Variants, PackageType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftPackage(name, Options, Variants, PackageType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftPackage']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftPackage(params['id']).subscribe(res => {
                this.aircraftPackage = res;
            });
        });
    }
}