
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWarrantyComponent } from './create.component';
import { WarrantyService } from '../../../services/Warranty.service';
import { Router } from '@angular/router';

describe('CreateWarrantyComponent', () => {
  let component: CreateWarrantyComponent;
  let fixture: ComponentFixture<CreateWarrantyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWarrantyComponent
      ],
      providers: [
        WarrantyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWarrantyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});