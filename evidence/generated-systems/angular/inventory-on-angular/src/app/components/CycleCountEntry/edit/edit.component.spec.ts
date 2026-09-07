
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditCycleCountEntryComponent } from './edit.component';
import { CycleCountEntryService } from '../../../services/CycleCountEntry.service';

describe('EditCycleCountEntryComponent', () => {
  let component: EditCycleCountEntryComponent;
  let fixture: ComponentFixture<EditCycleCountEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditCycleCountEntryComponent
      ],
      providers: [
        CycleCountEntryService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditCycleCountEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});