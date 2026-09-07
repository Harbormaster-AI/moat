
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLandingGearComponent } from './create.component';
import { LandingGearService } from '../../../services/LandingGear.service';
import { Router } from '@angular/router';

describe('CreateLandingGearComponent', () => {
  let component: CreateLandingGearComponent;
  let fixture: ComponentFixture<CreateLandingGearComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLandingGearComponent
      ],
      providers: [
        LandingGearService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLandingGearComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});