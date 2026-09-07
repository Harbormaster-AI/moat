
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftVariantComponent } from './create.component';
import { AircraftVariantService } from '../../../services/AircraftVariant.service';
import { Router } from '@angular/router';

describe('CreateAircraftVariantComponent', () => {
  let component: CreateAircraftVariantComponent;
  let fixture: ComponentFixture<CreateAircraftVariantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftVariantComponent
      ],
      providers: [
        AircraftVariantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftVariantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});