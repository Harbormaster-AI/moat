
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPriceBookEntryComponent } from './edit.component';
import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';

describe('EditPriceBookEntryComponent', () => {
  let component: EditPriceBookEntryComponent;
  let fixture: ComponentFixture<EditPriceBookEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPriceBookEntryComponent
      ],
      providers: [
        PriceBookEntryService,
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

    fixture = TestBed.createComponent(EditPriceBookEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});