
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateImagingOrderComponent } from './create.component';
import { ImagingOrderService } from '../../../services/ImagingOrder.service';
import { Router } from '@angular/router';

describe('CreateImagingOrderComponent', () => {
  let component: CreateImagingOrderComponent;
  let fixture: ComponentFixture<CreateImagingOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateImagingOrderComponent
      ],
      providers: [
        ImagingOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateImagingOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});