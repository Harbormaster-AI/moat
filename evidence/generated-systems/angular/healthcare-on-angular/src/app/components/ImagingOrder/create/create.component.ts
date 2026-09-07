import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ImagingOrderService } from '../../../services/ImagingOrder.service';
import { ImagingOrder } from '../../../models/ImagingOrder';
import { SubBaseComponent } from '../../ImagingOrder/sub.base.component';

@Component({
    selector: 'app-create-imagingOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateImagingOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add ImagingOrder';

    imagingOrderForm: FormGroup;
    imagingOrder: ImagingOrder;

    constructor( http: HttpClient,
        private imagingOrderService: ImagingOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.imagingOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  bodySite: ['', Validators.required],
      contrast: ['', Validators.required],
      Order: ['', ],
      ImagingCenter: ['', ],
      Reports: ['', ],
      Modality: ['', ]
        });
    }

    
    addImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality): void {
        this.imagingOrderService
        .addImagingOrder(bodySite, contrast, Order, ImagingCenter, Reports, Modality)
            .subscribe(() => {
                this.router.navigate(['/indexImagingOrder']);
            });
    }

    ngOnInit(): void {
    }
}