
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDemandSignalComponent } from './create.component';
import { DemandSignalService } from '../../../services/DemandSignal.service';
import { Router } from '@angular/router';

describe('CreateDemandSignalComponent', () => {
  let component: CreateDemandSignalComponent;
  let fixture: ComponentFixture<CreateDemandSignalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDemandSignalComponent
      ],
      providers: [
        DemandSignalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDemandSignalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});