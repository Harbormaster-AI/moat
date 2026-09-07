import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftProgramService } from '../../../services/AircraftProgram.service';
import { SubBaseComponent } from '../../AircraftProgram/sub.base.component';


@Component({
    selector: 'app-edit-aircraftProgram',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftProgram';

    aircraftProgramForm: FormGroup;
    aircraftProgram: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftProgramService,
        private fb: FormBuilder
) {
        super(http);
        this.aircraftProgramForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      programCode: ['', Validators.required],
      entryIntoServiceYear: ['', Validators.required],
      Manufacturer: ['', ],
      AircraftFamilies: ['', ],
      TypeCertificate: ['', ],
      KeySuppliers: ['', ],
      Status: ['', ]
        });
    }

    
    updateAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftProgram']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftProgram(params['id']).subscribe(res => {
                this.aircraftProgram = res;
            });
        });
    }
}