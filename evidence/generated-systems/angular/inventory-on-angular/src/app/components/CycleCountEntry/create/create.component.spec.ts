
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCycleCountEntryComponent } from './create.component';
import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';
import { Router } from '@angular/router';

describe('CreateCycleCountEntryComponent', () => {
  let component: CreateCycleCountEntryComponent;
  let fixture: ComponentFixture<CreateCycleCountEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCycleCountEntryComponent
      ],
      providers: [
        CycleCountEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCycleCountEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});