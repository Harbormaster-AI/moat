import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BOMService } from '../../../services/BOM.service';
import { BOM } from '../../../models/BOM';
import { SubBaseComponent } from '../../BOM/sub.base.component';

@Component({
    selector: 'app-create-bOM',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBOMComponent extends SubBaseComponent implements OnInit {

    title = 'Add BOM';

    bOMForm: FormGroup;
    bOM: BOM;

    constructor( http: HttpClient,
        private bOMService: BOMService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.bOMForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  bomNumber: ['', Validators.required],
      revision: ['', Validators.required],
      effectivityStart: ['', Validators.required],
      effectivityEnd: ['', Validators.required],
      ParentItem: ['', ],
      BomItems: ['', ],
      Status: ['', ]
        });
    }

    
    addBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status): void {
        this.bOMService
        .addBOM(bomNumber, revision, effectivityStart, effectivityEnd, ParentItem, BomItems, Status)
            .subscribe(() => {
                this.router.navigate(['/indexBOM']);
            });
    }

    ngOnInit(): void {
    }
}