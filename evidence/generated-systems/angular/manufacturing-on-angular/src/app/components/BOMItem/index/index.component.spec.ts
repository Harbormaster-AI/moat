
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBOMItemComponent } from './index.component';
import { BOMItemService } from '../../../services/BOMItem.service';

describe('IndexBOMItemComponent', () => {
  let component: IndexBOMItemComponent;
  let fixture: ComponentFixture<IndexBOMItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBOMItemComponent
      ],
      providers: [
        BOMItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBOMItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});