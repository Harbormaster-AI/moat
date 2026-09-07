
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEncounterComponent } from './create.component';
import { EncounterService } from '../../../services/Encounter.service';
import { Router } from '@angular/router';

describe('CreateEncounterComponent', () => {
  let component: CreateEncounterComponent;
  let fixture: ComponentFixture<CreateEncounterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEncounterComponent
      ],
      providers: [
        EncounterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEncounterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});