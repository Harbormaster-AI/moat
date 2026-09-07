
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditStockKeepingUnitComponent } from './edit.component';
import { StockKeepingUnitService } from '../../../services/StockKeepingUnit.service';

describe('EditStockKeepingUnitComponent', () => {
  let component: EditStockKeepingUnitComponent;
  let fixture: ComponentFixture<EditStockKeepingUnitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditStockKeepingUnitComponent
      ],
      providers: [
        StockKeepingUnitService,
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

    fixture = TestBed.createComponent(EditStockKeepingUnitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});