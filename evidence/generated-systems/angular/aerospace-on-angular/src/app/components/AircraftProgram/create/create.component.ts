import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AircraftProgramService } from '../../../services/AircraftProgram.service';
import { AircraftProgram } from '../../../models/AircraftProgram';
import { SubBaseComponent } from '../../AircraftProgram/sub.base.component';

@Component({
    selector: 'app-create-aircraftProgram',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAircraftProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Add AircraftProgram';

    aircraftProgramForm: FormGroup;
    aircraftProgram: AircraftProgram;

    constructor( http: HttpClient,
        private aircraftProgramService: AircraftProgramService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status): void {
        this.aircraftProgramService
        .addAircraftProgram(name, programCode, entryIntoServiceYear, Manufacturer, AircraftFamilies, TypeCertificate, KeySuppliers, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAircraftProgram']);
            });
    }

    ngOnInit(): void {
    }
}