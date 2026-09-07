
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftProgramComponent } from './create.component';
import { AircraftProgramService } from '../../../services/AircraftProgram.service';
import { Router } from '@angular/router';

describe('CreateAircraftProgramComponent', () => {
  let component: CreateAircraftProgramComponent;
  let fixture: ComponentFixture<CreateAircraftProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftProgramComponent
      ],
      providers: [
        AircraftProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftProgramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});