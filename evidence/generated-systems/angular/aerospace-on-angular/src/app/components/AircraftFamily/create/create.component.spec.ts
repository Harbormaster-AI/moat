
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftFamilyComponent } from './create.component';
import { AircraftFamilyService } from '../../../services/AircraftFamily.service';
import { Router } from '@angular/router';

describe('CreateAircraftFamilyComponent', () => {
  let component: CreateAircraftFamilyComponent;
  let fixture: ComponentFixture<CreateAircraftFamilyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftFamilyComponent
      ],
      providers: [
        AircraftFamilyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftFamilyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});