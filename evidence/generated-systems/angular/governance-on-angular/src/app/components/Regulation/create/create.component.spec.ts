
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRegulationComponent } from './create.component';
import { RegulationService } from '../../../services/Regulation.service';
import { Router } from '@angular/router';

describe('CreateRegulationComponent', () => {
  let component: CreateRegulationComponent;
  let fixture: ComponentFixture<CreateRegulationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRegulationComponent
      ],
      providers: [
        RegulationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRegulationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});