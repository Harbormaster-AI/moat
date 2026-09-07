
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProductOfferingComponent } from './create.component';
import { ProductOfferingService } from '../../../services/ProductOffering.service';
import { Router } from '@angular/router';

describe('CreateProductOfferingComponent', () => {
  let component: CreateProductOfferingComponent;
  let fixture: ComponentFixture<CreateProductOfferingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProductOfferingComponent
      ],
      providers: [
        ProductOfferingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProductOfferingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});